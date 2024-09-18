package response

type HubSession struct {
	NameStr               string `json:"Name_str"`
	RemoteSessionBool     bool   `json:"RemoteSession_bool"`
	RemoteHostnameStr     string `json:"RemoteHostname_str"`
	UsernameStr           string `json:"Username_str"`
	ClientIp              string `json:"ClientIP_ip"`
	HostnameStr           string `json:"Hostname_str"`
	MaxNumTcpU32          int    `json:"MaxNumTcp_u32"`
	CurrentNumTcpU32      int    `json:"CurrentNumTcp_u32"`
	PacketSizeU64         int    `json:"PacketSize_u64"`
	PacketNumU64          int    `json:"PacketNum_u64"`
	LinkModeBool          bool   `json:"LinkMode_bool"`
	SecureNATModeBool     bool   `json:"SecureNATMode_bool"`
	BridgeModeBool        bool   `json:"BridgeMode_bool"`
	Layer3ModeBool        bool   `json:"Layer3Mode_bool"`
	ClientBridgeModeBool  bool   `json:"Client_BridgeMode_bool"`
	ClientMonitorModeBool bool   `json:"Client_MonitorMode_bool"`
	VLanIdU32             int    `json:"VLanId_u32"`
	UniqueIdBin           string `json:"UniqueId_bin"`
	CreatedTimeDt         string `json:"CreatedTime_dt"`
	LastCommTimeDt        string `json:"LastCommTime_dt"`
}
type ResHubSessionList struct {
	HubNameStr  string        `json:"HubName_str"`
	SessionList []*HubSession `json:"SessionList"`
}

type Session struct {
	Account     string `json:"account"`
	SessionName string `json:"session_name"`
	ClientIp    string `json:"client_ip"`
	ClientMac   string `json:"client_mac"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}
