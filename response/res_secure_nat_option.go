package response

type ResSecureNATOption struct {
	RpcHubNameStr           string `json:"RpcHubName_str"`
	MacAddressBin           string `json:"MacAddress_bin"`
	IpIp                    string `json:"Ip_ip"`
	MaskIp                  string `json:"Mask_ip"`
	UseNatBool              bool   `json:"UseNat_bool"`
	MtuU32                  int    `json:"Mtu_u32"`
	NatTcpTimeoutU32        int    `json:"NatTcpTimeout_u32"`
	NatUdpTimeoutU32        int    `json:"NatUdpTimeout_u32"`
	UseDhcpBool             bool   `json:"UseDhcp_bool"`
	DhcpLeaseIPStartIp      string `json:"DhcpLeaseIPStart_ip"`
	DhcpLeaseIPEndIp        string `json:"DhcpLeaseIPEnd_ip"`
	DhcpSubnetMaskIp        string `json:"DhcpSubnetMask_ip"`
	DhcpExpireTimeSpanU32   int    `json:"DhcpExpireTimeSpan_u32"`
	DhcpGatewayAddressIp    string `json:"DhcpGatewayAddress_ip"`
	DhcpDnsServerAddressIp  string `json:"DhcpDnsServerAddress_ip"`
	DhcpDnsServerAddress2Ip string `json:"DhcpDnsServerAddress2_ip"`
	DhcpDomainNameStr       string `json:"DhcpDomainName_str"`
	SaveLogBool             bool   `json:"SaveLog_bool"`
	ApplyDhcpPushRoutesBool bool   `json:"ApplyDhcpPushRoutes_bool"`
	DhcpPushRoutesStr       string `json:"DhcpPushRoutes_str"`
}
