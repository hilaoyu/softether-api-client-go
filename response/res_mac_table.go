package response

type MacTable struct {
	KeyU32            int    `json:"Key_u32"`
	SessionNameStr    string `json:"SessionName_str"`
	MacAddressBin     string `json:"MacAddress_bin"`
	CreatedTimeDt     string `json:"CreatedTime_dt"`
	UpdatedTimeDt     string `json:"UpdatedTime_dt"`
	RemoteItemBool    bool   `json:"RemoteItem_bool"`
	RemoteHostnameStr string `json:"RemoteHostname_str"`
	VlanIdU32         int    `json:"VlanId_u32"`
}
type ResMacTableList struct {
	HubNameStr string       `json:"HubName_str"`
	MacTable   *[]*MacTable `json:"MacTable"`
}
