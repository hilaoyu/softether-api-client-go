package response

type ResResult struct {
	JsonRpc string      `json:"jsonrpc"`
	Id      string      `json:"id"`
	Result  interface{} `json:"result"`
}
