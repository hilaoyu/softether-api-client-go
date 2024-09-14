package softetherapi

import (
	"encoding/json"
	"fmt"
	"github.com/hilaoyu/go-utils/utilHttp"
	"github.com/hilaoyu/softether-api-client-go/methods"
	"github.com/hilaoyu/softether-api-client-go/pkg"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// to handle softether vpn server json-rpc api
type Api struct {
	HttpClient *utilHttp.HttpClient
	Hub        string
	Password   string
	id         Id
}

// Id interface is used for request id
type Id interface {
	Incl()
	Describe() int
}

// id struct is the implement of Id interface
type id int

func initId() Id {
	return id(0)
}

// This method increment an id
func (i id) Incl() {
	i += 1
}

func (i id) Describe() int {
	return int(i)
}

func New(host string, port int, hub, password string) *Api {
	httpClient := utilHttp.NewHttpClient(fmt.Sprintf("https://%s:%d/api", host, port), time.Duration(10)*time.Second)
	httpClient.SetSslVerify(false)
	return &Api{
		HttpClient: httpClient,
		Hub:        hub,
		Password:   password,
		id:         initId(),
	}
}

func (api *Api) UseProxySocks5(proxyAddr string, proxyUser string, proxyPassword string) *Api {
	api.HttpClient = api.HttpClient.UseProxySocks5(proxyAddr, proxyUser, proxyPassword)
	return api
}

func (api *Api) Call(method pkg.Method) (map[string]interface{}, error) {
	api.id.Incl()
	method.SetId(api.id.Describe())
	body, err := method.Marshall()
	//fmt.Println(string(body))
	if err != nil {
		return nil, fmt.Errorf("[error] failed to marshall request: %v", err)
	}

	res, err := api.HttpClient.WithRawBody(string(body)).Request("POST", "", map[string]string{
		"Content-Type":        "application/json",
		"X-VPNADMIN-HUBNAME":  api.Hub,
		"X-VPNADMIN-PASSWORD": api.Password,
	})

	if err != nil {
		return nil, fmt.Errorf("[error] failed to call: %v", err)
	}
	return validateResponse(res)
}

func validateResponse(res *http.Response) (map[string]interface{}, error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("[error] failed to read: %v", err)
	}
	//fmt.Println(string(body))
	b := make(map[string]interface{})
	err = json.Unmarshal(body, &b)
	if err != nil {
		return nil, err
	}
	if _, ok := b["result"]; ok {
		return b["result"].(map[string]interface{}), nil
	}
	return nil, fmt.Errorf("%v", b["error"])
}

func (api *Api) ServerInfo() (info map[string]interface{}, err error) {
	info, err = api.Call(methods.NewTest())
	return
}

func (api *Api) HubCreate(hubName string, password string, natOption *methods.SetSecureNATOptionParams) (err error) {
	if "" == hubName {
		err = fmt.Errorf("hub name can not be null")
		return
	}

	api.Call(methods.NewDeleteHub(hubName))

	_, err = api.Call(methods.NewCreateHub(hubName, password, true))
	if nil != err {
		err = fmt.Errorf("创建HUB失败, err: %+v", err)
		return
	}

	if nil == natOption {
		return
	}
	_, err = api.Call(methods.NewEnableSecureNAT(hubName))
	if nil != err {
		err = fmt.Errorf("HUB: %s 启用NAT失败,err: %+v", hubName, err)
		return
	}

	hubOption, err := api.Call(methods.NewGetSecureNATOption(hubName))
	if nil != err {
		err = fmt.Errorf("HUB: %s 查询 SecureNATOption 失败", hubName)
	}

	macBin := ""
	if macBinI, ok := hubOption["MacAddress_bin"]; ok {
		macBin = macBinI.(string)
	}
	natOption.MacAddressBin = macBin
	natOption.RpcHubName = hubName
	if natOption.Mtu <= 0 {
		natOption.Mtu = 1500
	}
	if natOption.NatTcpTimeout <= 0 {
		natOption.NatTcpTimeout = 1800
	}
	if natOption.NatUdpTimeout <= 0 {
		natOption.NatUdpTimeout = 120
	}
	if natOption.DhcpExpireTimeSpan <= 0 {
		natOption.DhcpExpireTimeSpan = 7200
	}
	secureNatOptionMethod := methods.NewSetSecureNATOption(hubName)
	secureNatOptionMethod.Params = natOption

	_, err = api.Call(secureNatOptionMethod)
	if nil != err {
		err = fmt.Errorf("HUB: %s 设置NAT OPTIONS失败,%+v", hubName, err)
	}
	return
}

func (api *Api) HubDelete(hubName string) (err error) {
	_, err = api.Call(methods.NewDeleteHub(hubName))
	return
}

func (api *Api) HubLinkCreate(hubName string, linkName string, destHost string, destPort int, destHubName string, account string, password string, onLine bool) (err error) {

	m := methods.NewCreateLink(hubName, linkName)

	m.Params.Hostname = destHost
	m.Params.Port = destPort
	m.Params.HubName = destHubName
	m.Params.Username = account
	m.Params.PlainPassword = password
	m.Params.AuthType = 2
	_, err = api.Call(m)
	if nil != err {
		return
	}

	if onLine {
		_, err = api.Call(methods.NewSetLinkOnline(hubName, linkName))
	} else {
		_, err = api.Call(methods.NewSetLinkOffline(hubName, linkName))
	}

	return
}

func (api *Api) HubLinkDelete(hubName string, linkName string) (err error) {
	_, err = api.Call(methods.NewDeleteLink(hubName, linkName))

	return
}

func (api *Api) AccountCreate(hubName string, account string, password string) (err error) {

	m := methods.NewCreateUser(hubName, account, account, "", nil, 1)
	m.Params.Auth_Passoword = password

	_, err = api.Call(m)

	return
}

func (api *Api) AccountDelete(hubName string, account string) (err error) {

	m := methods.NewDeleteUser(hubName, account)
	_, err = api.Call(m)
	return
}

func ConvStringPort(str string) (port int) {
	port, err := strconv.Atoi(str)
	if nil != err {
		port = 5555
	}
	return
}
