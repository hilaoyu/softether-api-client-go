package response

type ResCreateUser struct {
	HubNameStr                                string `json:"HubName_str"`
	NameStr                                   string `json:"Name_str"`
	GroupNameStr                              string `json:"GroupName_str"`
	RealnameUtf                               string `json:"Realname_utf"`
	NoteUtf                                   string `json:"Note_utf"`
	CreatedTimeDt                             string `json:"CreatedTime_dt"`
	UpdatedTimeDt                             string `json:"UpdatedTime_dt"`
	ExpireTimeDt                              string `json:"ExpireTime_dt"`
	AuthTypeU32                               int    `json:"AuthType_u32"`
	AuthPasswordStr                           string `json:"Auth_Password_str"`
	UserXBin                                  string `json:"UserX_bin"`
	SerialBin                                 string `json:"Serial_bin"`
	CommonNameUtf                             string `json:"CommonName_utf"`
	RadiusUsernameUtf                         string `json:"RadiusUsername_utf"`
	NtUsernameUtf                             string `json:"NtUsername_utf"`
	NumLoginU32                               int    `json:"NumLogin_u32"`
	RecvBroadcastBytesU64                     int    `json:"Recv.BroadcastBytes_u64"`
	RecvBroadcastCountU64                     int    `json:"Recv.BroadcastCount_u64"`
	RecvUnicastBytesU64                       int    `json:"Recv.UnicastBytes_u64"`
	RecvUnicastCountU64                       int    `json:"Recv.UnicastCount_u64"`
	SendBroadcastBytesU64                     int    `json:"Send.BroadcastBytes_u64"`
	SendBroadcastCountU64                     int    `json:"Send.BroadcastCount_u64"`
	SendUnicastBytesU64                       int    `json:"Send.UnicastBytes_u64"`
	SendUnicastCountU64                       int    `json:"Send.UnicastCount_u64"`
	UsePolicyBool                             bool   `json:"UsePolicy_bool"`
	PolicyAccessBool                          bool   `json:"policy:Access_bool"`
	PolicyDHCPFilterBool                      bool   `json:"policy:DHCPFilter_bool"`
	PolicyDHCPNoServerBool                    bool   `json:"policy:DHCPNoServer_bool"`
	PolicyDHCPForceBool                       bool   `json:"policy:DHCPForce_bool"`
	PolicyNoBridgeBool                        bool   `json:"policy:NoBridge_bool"`
	PolicyNoRoutingBool                       bool   `json:"policy:NoRouting_bool"`
	PolicyCheckMacBool                        bool   `json:"policy:CheckMac_bool"`
	PolicyCheckIPBool                         bool   `json:"policy:CheckIP_bool"`
	PolicyArpDhcpOnlyBool                     bool   `json:"policy:ArpDhcpOnly_bool"`
	PolicyPrivacyFilterBool                   bool   `json:"policy:PrivacyFilter_bool"`
	PolicyNoServerBool                        bool   `json:"policy:NoServer_bool"`
	PolicyNoBroadcastLimiterBool              bool   `json:"policy:NoBroadcastLimiter_bool"`
	PolicyMonitorPortBool                     bool   `json:"policy:MonitorPort_bool"`
	PolicyMaxConnectionU32                    int    `json:"policy:MaxConnection_u32"`
	PolicyTimeOutU32                          int    `json:"policy:TimeOut_u32"`
	PolicyMaxMacU32                           int    `json:"policy:MaxMac_u32"`
	PolicyMaxIPU32                            int    `json:"policy:MaxIP_u32"`
	PolicyMaxUploadU32                        int    `json:"policy:MaxUpload_u32"`
	PolicyMaxDownloadU32                      int    `json:"policy:MaxDownload_u32"`
	PolicyFixPasswordBool                     bool   `json:"policy:FixPassword_bool"`
	PolicyMultiLoginsU32                      int    `json:"policy:MultiLogins_u32"`
	PolicyNoQoSBool                           bool   `json:"policy:NoQoS_bool"`
	PolicyRSandRAFilterBool                   bool   `json:"policy:RSandRAFilter_bool"`
	PolicyRAFilterBool                        bool   `json:"policy:RAFilter_bool"`
	PolicyDHCPv6FilterBool                    bool   `json:"policy:DHCPv6Filter_bool"`
	PolicyDHCPv6NoServerBool                  bool   `json:"policy:DHCPv6NoServer_bool"`
	PolicyNoRoutingV6Bool                     bool   `json:"policy:NoRoutingV6_bool"`
	PolicyCheckIPv6Bool                       bool   `json:"policy:CheckIPv6_bool"`
	PolicyNoServerV6Bool                      bool   `json:"policy:NoServerV6_bool"`
	PolicyMaxIPv6U32                          int    `json:"policy:MaxIPv6_u32"`
	PolicyNoSavePasswordBool                  bool   `json:"policy:NoSavePassword_bool"`
	PolicyAutoDisconnectU32                   int    `json:"policy:AutoDisconnect_u32"`
	PolicyFilterIPv4Bool                      bool   `json:"policy:FilterIPv4_bool"`
	PolicyFilterIPv6Bool                      bool   `json:"policy:FilterIPv6_bool"`
	PolicyFilterNonIPBool                     bool   `json:"policy:FilterNonIP_bool"`
	PolicyNoIPv6DefaultRouterInRABool         bool   `json:"policy:NoIPv6DefaultRouterInRA_bool"`
	PolicyNoIPv6DefaultRouterInRAWhenIPv6Bool bool   `json:"policy:NoIPv6DefaultRouterInRAWhenIPv6_bool"`
	PolicyVLanIdU32                           int    `json:"policy:VLanId_u32"`
	PolicyVer3Bool                            bool   `json:"policy:Ver3_bool"`
}
