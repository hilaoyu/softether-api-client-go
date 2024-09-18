package response

type ResServerInfo struct {
	ServerProductNameStr     string `json:"ServerProductName_str"`
	ServerVersionStringStr   string `json:"ServerVersionString_str"`
	ServerBuildInfoStringStr string `json:"ServerBuildInfoString_str"`
	ServerVerIntU32          int    `json:"ServerVerInt_u32"`
	ServerBuildIntU32        int    `json:"ServerBuildInt_u32"`
	ServerHostNameStr        string `json:"ServerHostName_str"`
	ServerTypeU32            int    `json:"ServerType_u32"`
	ServerBuildDateDt        string `json:"ServerBuildDate_dt"`
	ServerFamilyNameStr      string `json:"ServerFamilyName_str"`
	OsTypeU32                int    `json:"OsType_u32"`
	OsServicePackU32         int    `json:"OsServicePack_u32"`
	OsSystemNameStr          string `json:"OsSystemName_str"`
	OsProductNameStr         string `json:"OsProductName_str"`
	OsVendorNameStr          string `json:"OsVendorName_str"`
	OsVersionStr             string `json:"OsVersion_str"`
	KernelNameStr            string `json:"KernelName_str"`
	KernelVersionStr         string `json:"KernelVersion_str"`
}
