# Minecraft Resource for testing
Server:
- Java: j4.1919801.xyz (SRV) , mc6.1919801.xyz:11451
- Bedrock: mc6.1919801.xyz
User:
- wor114514
Data examples:
Under ./examples
Fonts:
.\icons_and_translates\fonts\*
# Skin API
For displaying player skin image on the web
Full skin: https://mineskin.eu/skin/{player}
2D Head: https://mineskin.eu/helm/{player}
# Get Server Status (mcutil)
For getting server status at backend
https://pkg.go.dev/github.com/mcstatus-io/mcutil
Useful:
- StatusBedrock
- Status
# Skinview 3D
For 3D display of skin on the web, you can get the skin image from Skin API(Full skin)
https://github.com/bs-community/skinview3d/
# GOPSUTIL
For getting hardware info at backend
https://github.com/shirou/gopsutil
Useful:
- PlatformVersion 
- VirtualizationSystem 
- virtual_memory
- cpu_percent
- disk_io_counters
- net_io_counters
- ModelName 
- Load1, Load5, Load15
# Minecraft names & icons
Blocks translation and icon name: .\icons_and_translates\objectives\blocks.xml
Blocks icons: .\icons_and_translates\sprites\blocks\*
Advancements translation and icon name: .\icons_and_translates\objectives\advancements\*
Advancement icons: .\icons_and_translates\sprites\global\advancements
Menu icons: .\icons_and_translates\icons
Items icons: .\icons_and_translates\sprites\global\items
Other icons: .\icons_and_translates\sprites\global\criteria
# Parsing .dat
https://github.com/Offroaders123/NBTify/
# Stun & DDNS
## Stun
http://{lucky_server_ip}:{lucky_server_port}/{lucky_safe_entrance}/api/stunrulelist?openToken={lucky_token}
Example Response:
{"ModuleEnable":true,"list":[{"Key":"hpXxd4se5Px2T3Gs","Name":"Jvav4","StunType":"tcp4","Enable":true,"DisablePortForward":false,"LastLogs":null,"StunLocalAddr":":43861","TargetAddrList":["127.0.0.1:11451"],"PublicAddr":"120.229.204.196:10778","PublicAddrInfo":"","PublicAddrHistroy":[{"AddrRecord":"120.229.204.196:18391","UpdateTime":"2026-04-13 00:49:30"},{"AddrRecord":"","UpdateTime":"2026-04-13 01:15:45"},{"AddrRecord":"120.229.204.196:10778","UpdateTime":"2026-04-13 01:15:50"}],"WebhookEnable":false,"WebhookProxy":"","WebhookCallTime":"","WebhookCallResult":false,"WebhookCallErrorMsg":"","WebhookCallHistroy":null,"GlobalWebhook":false,"GlobalWebhookCallTime":"","GlobalWebhookCallResult":false,"GlobalWebhookCallErrorMsg":"","GlobalWebhookCallHistroy":null,"Options":{"DisableSelfForwardingCheck":false,"SingleProxyMaxTCPConnections":256,"SingleProxyMaxUDPReadTargetDatagoroutineCount":32,"UDPSessionTimeout":30000,"SafeMode":"blacklist","TCPListenTLS":false,"TCPRelayTLS":false,"TCPRelayTLSServerName":"","TCPRelayTLSInsecureSkipVerify":false,"TCPStreamEncryptionSource":false,"TCPStreamEncryptionAccept":false,"TCPStreamEncryptionKey":"","SinglePortSpeedLimit":false,"SinglePortSendSpeedLimit":30,"SinglePortReceSpeedLimit":30,"RuleSpeedLimit":false,"RuleSendSpeedLimit":0,"RuleReceSpeedLimit":0,"UDPPacketSize":1500,"UDPPacketSourceEncryption":false,"UDPPacketAcceptEncryption":false,"UDPPacketEncryptionKey":""}},{"Key":"4Xv9n8GYa9bo07E3","Name":"BR4","StunType":"udp4","Enable":true,"DisablePortForward":false,"LastLogs":null,"StunLocalAddr":":58475","TargetAddrList":["127.0.0.1:11451"],"PublicAddr":"120.229.204.196:10830","PublicAddrInfo":"","PublicAddrHistroy":[{"AddrRecord":"120.229.204.196:18405","UpdateTime":"2026-04-13 00:49:32"},{"AddrRecord":"","UpdateTime":"2026-04-13 01:15:49"},{"AddrRecord":"120.229.204.196:10830","UpdateTime":"2026-04-13 01:15:54"}],"WebhookEnable":false,"WebhookProxy":"","WebhookCallTime":"","WebhookCallResult":false,"WebhookCallErrorMsg":"","WebhookCallHistroy":null,"GlobalWebhook":false,"GlobalWebhookCallTime":"","GlobalWebhookCallResult":false,"GlobalWebhookCallErrorMsg":"","GlobalWebhookCallHistroy":null,"Options":{"DisableSelfForwardingCheck":false,"SingleProxyMaxTCPConnections":256,"SingleProxyMaxUDPReadTargetDatagoroutineCount":32,"UDPSessionTimeout":30000,"SafeMode":"blacklist","TCPListenTLS":false,"TCPRelayTLS":false,"TCPRelayTLSServerName":"","TCPRelayTLSInsecureSkipVerify":false,"TCPStreamEncryptionSource":false,"TCPStreamEncryptionAccept":false,"TCPStreamEncryptionKey":"","SinglePortSpeedLimit":false,"SinglePortSendSpeedLimit":30,"SinglePortReceSpeedLimit":30,"RuleSpeedLimit":false,"RuleSendSpeedLimit":0,"RuleReceSpeedLimit":0,"UDPPacketSize":1500,"UDPPacketSourceEncryption":false,"UDPPacketAcceptEncryption":false,"UDPPacketEncryptionKey":""}}],"ret":0,"statistics":{"4Xv9n8GYa9bo07E3":{"TrafficIn":2727710,"TrafficOut":64395114,"TCPCurrentConnections":0,"UDPCurrentConnections":0},"hpXxd4se5Px2T3Gs":{"TrafficIn":62801244,"TrafficOut":4563872832,"TCPCurrentConnections":24,"UDPCurrentConnections":24}}}
## DDNS
http://{lucky_server_ip}:{lucky_server_port}/{lucky_safe_entrance}/api/ddnstasklist?openToken={lucky_token}
Example Response:
{"data":[{"TaskName":"CF","TaskKey":"XmVNUDdZq4MjgAkq","TaskType":"IPv4\u0026IPv6","Enable":true,"Expanded":true,"IPSectionExpanded":true,"TTL":"","DNS":{"Name":"cloudflare","HttpClientProxyType":"","HttpClientProxyAddr":"","Callback":{"Server":""}},"IPV4AddrHistory":[{"IPaddr":"120.229.204.196","RecordTime":"2026-04-12 23:52:37"}],"IPV6AddrHistory":[{"IPaddr":"2409:8a55:3191:7851::ea1","RecordTime":"2026-04-12 23:52:47"}],"Ipv4Addr":"120.229.204.196","Ipv6Addr":"2409:8a55:3191:7851::ea1","Ipv4AddrErrMsg":"","Ipv6AddrErrMsg":"","Ipv4AddrInfo":"","Ipv6AddrInfo":"","DebugMode":false,"WebhookEnable":false,"GlobalWebhook":false,"IngoreWebhookVariablesNotFound":false,"IngoreWebhookVariablesNotFoundList":"","WebhookProxy":"","WebhookCallTime":"","WebhookCallResult":false,"WebhookCallErrorMsg":"","WebhookCallHistroy":null,"GlobalWebhookCallTime":"","GlobalWebhookCallResult":false,"GlobalWebhookCallErrorMsg":"","GlobalWebhookCallHistroy":null,"SyncTimeHistroy":["2026-04-15 22:36:30","2026-04-15 22:37:06","2026-04-15 22:37:42","2026-04-15 22:38:19","2026-04-15 22:38:55","2026-04-15 22:39:31","2026-04-15 22:40:07","2026-04-15 22:40:43","2026-04-15 22:41:19","2026-04-15 22:41:55"],"Records":[{"Remark":"","DomainName":"1919801.xyz","SubDomain":"mc6","UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","LastUpdateStatusTime":"2026-04-15 22:41:56","Message":"","UpdateHistroy":[{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:36:30"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:37:06"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:37:43"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:38:19"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:38:55"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:39:31"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:40:07"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:40:43"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:41:19"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:41:56"}],"RecordHistory":[{"Records":[{"Value":"2409:8a55:3191:7851::ea1","Info":""}],"Time":"2026-04-12 23:52:52"}],"LastRecord":[{"Record":"2409:8a55:3191:7851::ea1","Info":""}],"Type":"AAAA","Key":"GNXJc741PwnZqjsb","Line":null,"Enable":true},{"Remark":"","DomainName":"1919801.xyz","SubDomain":"mc4","UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","LastUpdateStatusTime":"2026-04-15 22:41:56","Message":"","UpdateHistroy":[{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:36:30"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:37:06"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:37:43"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:38:19"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:38:55"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:39:31"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:40:07"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:40:43"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:41:19"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:41:56"}],"RecordHistory":[{"Records":[{"Value":"120.229.204.196","Info":""}],"Time":"2026-04-12 23:52:52"}],"LastRecord":[{"Record":"120.229.204.196","Info":""}],"Type":"A","Key":"1RrV4qOVnoS5inXG","Line":null,"Enable":true},{"Remark":"","DomainName":"1919801.xyz","SubDomain":"_minecraft._tcp.j4","UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","LastUpdateStatusTime":"2026-04-15 22:41:56","Message":"","UpdateHistroy":[{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:36:30"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:37:06"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:37:43"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:38:19"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:38:55"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:39:31"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:40:07"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:40:43"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:41:19"},{"UpdateStatus":"SYNC_LOC_RECORD_NOCHANGE","UpdateTime":"2026-04-15 22:41:56"}],"RecordHistory":[{"Records":[{"Value":"5 65535 18391 mc4.1919801.xyz.","Info":""}],"Time":"2026-04-13 00:49:38"},{"Records":[{"Value":"5 65535 18391 120.229.204.196.","Info":""}],"Time":"2026-04-13 01:14:45"},{"Records":[{"Value":"5 65535 18391 mc4.1919801.xyz.","Info":""}],"Time":"2026-04-13 01:15:02"},{"Records":[{"Value":"5 65535 10778 mc4.1919801.xyz.","Info":""}],"Time":"2026-04-13 01:15:55"}],"LastRecord":[{"Record":"5 65535 10778 mc4.1919801.xyz.","Info":""}],"Type":"SRV","Key":"ASgeOkXaV1jkP9FC","Line":null,"Enable":true}],"V6QueryIPEnable":true,"V6QueryIPType":"netInterface","V4QueryIPEnable":true,"V4QueryIPType":"url","LastSyncTime":"2026-04-15 22:41:55","NextSyncTime":"2026-04-15 22:42:32","QueryingIPv4Addr":false,"QueryingIPv6Addr":false}],"ret":0}

