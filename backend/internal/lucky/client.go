// Package lucky provides a client for the Lucky STUN/DDNS API.
package lucky

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/mtn-server/backend/config"
)

// StunRule represents a single STUN port mapping rule.
type StunRule struct {
	Key        string `json:"Key"`
	Name       string `json:"Name"`
	StunType   string `json:"StunType"`
	Enable     bool   `json:"Enable"`
	PublicAddr string `json:"PublicAddr"`
}

// StunResponse is the Lucky STUN API response.
type StunResponse struct {
	ModuleEnable bool       `json:"ModuleEnable"`
	List         []StunRule `json:"list"`
	Ret          int        `json:"ret"`
}

// DDNSRecord represents a single DNS record.
type DDNSRecord struct {
	DomainName string `json:"DomainName"`
	SubDomain  string `json:"SubDomain"`
	Type       string `json:"Type"`
	LastRecord []struct {
		Record string `json:"Record"`
	} `json:"LastRecord"`
}

// DDNSTask represents a DDNS task with its records.
type DDNSTask struct {
	TaskName string       `json:"TaskName"`
	Ipv4Addr string       `json:"Ipv4Addr"`
	Ipv6Addr string       `json:"Ipv6Addr"`
	Records  []DDNSRecord `json:"Records"`
}

// DDNSResponse is the Lucky DDNS API response.
type DDNSResponse struct {
	Data []DDNSTask `json:"data"`
	Ret  int        `json:"ret"`
}

// ConnectionInfo holds resolved server connection information.
type ConnectionInfo struct {
	JavaIPv4    *AddressInfo `json:"java_ipv4,omitempty"`
	BedrockIPv4 *AddressInfo `json:"bedrock_ipv4,omitempty"`
	JavaIPv6    *AddressInfo `json:"java_ipv6,omitempty"`
	BedrockIPv6 *AddressInfo `json:"bedrock_ipv6,omitempty"`
	JavaSRV     *AddressInfo `json:"java_srv,omitempty"`
}

// AddressInfo holds a single address entry.
type AddressInfo struct {
	Domain string `json:"domain,omitempty"`
	IP     string `json:"ip"`
	Port   string `json:"port,omitempty"`
	Type   string `json:"type"` // "IPv4", "IPv6", "SRV"
}

// Client queries Lucky's STUN and DDNS APIs.
type Client struct {
	httpClient *http.Client
	baseURL    string
	token      string
	addresses  config.AddressesConfig

	mu       sync.RWMutex
	connInfo *ConnectionInfo
}

// New creates a new Lucky API client.
func New(cfg config.LuckyConfig, addresses config.AddressesConfig) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    cfg.ServerURL,
		token:      cfg.Token,
		addresses:  addresses,
	}
}

// Refresh fetches the latest STUN and DDNS data and builds connection info.
func (c *Client) Refresh(ctx context.Context) error {
	stunData, err := c.fetchStun(ctx)
	if err != nil {
		return fmt.Errorf("fetch stun: %w", err)
	}

	ddnsData, err := c.fetchDDNS(ctx)
	if err != nil {
		return fmt.Errorf("fetch ddns: %w", err)
	}

	info := c.buildConnectionInfo(stunData, ddnsData)

	c.mu.Lock()
	c.connInfo = info
	c.mu.Unlock()

	return nil
}

// GetConnectionInfo returns the latest connection information.
func (c *Client) GetConnectionInfo() *ConnectionInfo {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.connInfo == nil {
		return &ConnectionInfo{}
	}
	return c.connInfo
}

func (c *Client) fetchStun(ctx context.Context) (*StunResponse, error) {
	url := fmt.Sprintf("%s/api/stunrulelist?openToken=%s", c.baseURL, c.token)
	return fetchJSON[StunResponse](ctx, c.httpClient, url)
}

func (c *Client) fetchDDNS(ctx context.Context) (*DDNSResponse, error) {
	url := fmt.Sprintf("%s/api/ddnstasklist?openToken=%s", c.baseURL, c.token)
	return fetchJSON[DDNSResponse](ctx, c.httpClient, url)
}

func (c *Client) buildConnectionInfo(stun *StunResponse, ddns *DDNSResponse) *ConnectionInfo {
	info := &ConnectionInfo{}

	// Static IPv6 addresses from config
	info.JavaIPv6 = &AddressInfo{
		Domain: c.addresses.JavaIPv6,
		IP:     c.addresses.JavaIPv6,
		Type:   "IPv6",
	}
	info.BedrockIPv6 = &AddressInfo{
		Domain: c.addresses.BedrockIPv6,
		IP:     c.addresses.BedrockIPv6,
		Type:   "IPv6",
	}
	info.JavaSRV = &AddressInfo{
		Domain: c.addresses.JavaIPv4SRV,
		IP:     c.addresses.JavaIPv4SRV,
		Type:   "SRV",
	}

	// Build domain map from DDNS
	domainMap := make(map[string]string) // ip -> domain
	if ddns != nil {
		for _, task := range ddns.Data {
			for _, rec := range task.Records {
				if len(rec.LastRecord) > 0 {
					domain := rec.SubDomain + "." + rec.DomainName
					ip := rec.LastRecord[0].Record
					domainMap[ip] = domain
				}
			}
		}
	}

	// Match STUN rules to find IPv4 addresses
	if stun != nil {
		for _, rule := range stun.List {
			if !rule.Enable || rule.PublicAddr == "" {
				continue
			}
			addr := rule.PublicAddr
			// Try to find domain for this IP
			domain := ""
			for ip, d := range domainMap {
				if ip != "" && len(addr) > 0 {
					// Simple substring match for IP part
					if containsIP(addr, ip) {
						domain = d
						break
					}
				}
			}

			addrInfo := &AddressInfo{
				Domain: domain,
				IP:     addr,
				Type:   "IPv4",
			}

			if rule.StunType == "tcp4" {
				info.JavaIPv4 = addrInfo
			} else if rule.StunType == "udp4" {
				info.BedrockIPv4 = addrInfo
			}
		}
	}

	return info
}

func containsIP(addr, ip string) bool {
	// addr format: "1.2.3.4:port", ip format: "1.2.3.4"
	for i := 0; i < len(addr); i++ {
		if addr[i] == ':' {
			return addr[:i] == ip
		}
	}
	return addr == ip
}

func fetchJSON[T any](ctx context.Context, client *http.Client, url string) (*T, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return &result, nil
}
