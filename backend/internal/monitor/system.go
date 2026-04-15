package monitor

import (
	"runtime"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
)

// SystemStats holds server hardware and OS information.
type SystemStats struct {
	Platform       string  `json:"platform"`
	Virtualization string  `json:"virtualization"`
	CPUModel       string  `json:"cpu_model"`
	CPUPercent     float64 `json:"cpu_percent"`
	MemTotal       uint64  `json:"mem_total"`
	MemUsed        uint64  `json:"mem_used"`
	MemPercent     float64 `json:"mem_percent"`
	DiskRead       uint64  `json:"disk_read"`
	DiskWrite      uint64  `json:"disk_write"`
	NetSent        uint64  `json:"net_sent"`
	NetRecv        uint64  `json:"net_recv"`
	Load1          float64 `json:"load_1"`
	Load5          float64 `json:"load_5"`
	Load15         float64 `json:"load_15"`
}

func getSystemStats() *SystemStats {
	stats := &SystemStats{}

	// Platform info
	if info, err := host.Info(); err == nil {
		stats.Platform = info.Platform + " " + info.PlatformVersion
		stats.Virtualization = info.VirtualizationSystem
	}

	// CPU
	if percent, err := cpu.Percent(0, false); err == nil && len(percent) > 0 {
		stats.CPUPercent = percent[0]
	}
	if infos, err := cpu.Info(); err == nil && len(infos) > 0 {
		stats.CPUModel = infos[0].ModelName
	}

	// Memory
	if v, err := mem.VirtualMemory(); err == nil {
		stats.MemTotal = v.Total
		stats.MemUsed = v.Used
		stats.MemPercent = v.UsedPercent
	}

	// Disk I/O
	if counters, err := disk.IOCounters(); err == nil {
		for _, c := range counters {
			stats.DiskRead += c.ReadBytes
			stats.DiskWrite += c.WriteBytes
		}
	}

	// Network I/O
	if counters, err := net.IOCounters(false); err == nil && len(counters) > 0 {
		stats.NetSent = counters[0].BytesSent
		stats.NetRecv = counters[0].BytesRecv
	}

	// Load average (Linux only)
	if runtime.GOOS == "linux" {
		if avg, err := load.Avg(); err == nil {
			stats.Load1 = avg.Load1
			stats.Load5 = avg.Load5
			stats.Load15 = avg.Load15
		}
	}

	return stats
}
