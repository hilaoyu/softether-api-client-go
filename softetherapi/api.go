package softetherapi

import (
	"fmt"
	"github.com/hilaoyu/go-utils/utilHttp"
	"github.com/hilaoyu/go-utils/utilTime"
	"github.com/hilaoyu/softether-api-client-go/methods"
	"github.com/hilaoyu/softether-api-client-go/pkg"
	"github.com/hilaoyu/softether-api-client-go/response"
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

func (api *Api) Call(method pkg.Method, v interface{}) (err error) {
	api.id.Incl()
	method.SetId(api.id.Describe())
	body, err := method.Marshall()
	//fmt.Println(string(body))
	if err != nil {
		err = fmt.Errorf("[error] failed to marshall request: %v", err)
		return
	}

	result := &response.ResResult{
		Result: v,
	}

	err = api.HttpClient.WithRawBody(string(body)).RequestJson(result, "POST", "", map[string]string{
		"Content-Type":        "application/json",
		"X-VPNADMIN-HUBNAME":  api.Hub,
		"X-VPNADMIN-PASSWORD": api.Password,
	})

	if err != nil {
		err = fmt.Errorf("[error] failed to call: %v", err)
		return
	}
	v = result.Result
	return
}

func (api *Api) ServerInfo() (info *response.ResServerInfo, err error) {
	info = &response.ResServerInfo{}
	err = api.Call(methods.NewGetServerInfo(), info)
	if nil != err {
		return
	}
	if "" == info.KernelNameStr {
		info = nil
	}
	return
}

func (api *Api) HubCreate(hubName string, password string, natOption *methods.SetSecureNATOptionParams) (err error) {
	if "" == hubName {
		err = fmt.Errorf("hub name can not be null")
		return
	}

	api.Call(methods.NewDeleteHub(hubName), nil)

	err = api.Call(methods.NewCreateHub(hubName, password, true), nil)
	if nil != err {
		err = fmt.Errorf("创建HUB失败, err: %+v", err)
		return
	}

	if nil == natOption {
		err = api.Call(methods.NewDisableSecureNAT(hubName), nil)
		if nil != err {
			err = fmt.Errorf("HUB: %s 禁用NAT失败,err: %+v", hubName, err)
		}
		return
	}
	err = api.Call(methods.NewEnableSecureNAT(hubName), nil)
	if nil != err {
		err = fmt.Errorf("HUB: %s 启用NAT失败,err: %+v", hubName, err)
		return
	}
	hubOption, err := api.HubGetNatOption(hubName)
	if nil != err {
		return
	}

	if nil == hubOption {
		err = fmt.Errorf("HUB: %s NAT错误", hubName)
		return
	}
	if "" != hubOption.MacAddressBin {
		natOption.MacAddressBin = hubOption.MacAddressBin
	}

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

	err = api.Call(secureNatOptionMethod, nil)
	if nil != err {
		err = fmt.Errorf("HUB: %s 设置NAT OPTIONS失败,%+v", hubName, err)
	}
	return
}

func (api *Api) HubGetNatOption(hubName string) (hubOption *response.ResSecureNATOption, err error) {
	hubOption = &response.ResSecureNATOption{}
	err = api.Call(methods.NewGetSecureNATOption(hubName), hubOption)
	if nil != err {
		err = fmt.Errorf("HUB: %s 查询 SecureNATOption 失败", hubName)
	}
	if "" == hubOption.RpcHubNameStr {
		hubOption = nil
	}
	return
}
func (api *Api) HubSessions(hubName string) (sessions []*response.Session, err error) {
	hubSessionList := &response.ResHubSessionList{}
	err = api.Call(methods.NewEnumSession(hubName), hubSessionList)
	if nil != err {
		err = fmt.Errorf("HUB: %s 查询 sessions 失败", hubName)
		return
	}

	if len(hubSessionList.SessionList) > 0 {
		var macTables []*response.MacTable
		macTableList := &response.ResMacTableList{MacTable: &macTables}
		err = api.Call(methods.NewEnumMacTable(hubName), macTableList)
		if nil != err {
			err = fmt.Errorf("HUB: %s 查询 mac table 失败", hubName)
			return
		}
		macTablesMap := map[string]*response.MacTable{}
		for _, macTable := range macTables {
			macTablesMap[macTable.SessionNameStr] = macTable
		}

		for _, hubSession := range hubSessionList.SessionList {
			if "" == hubSession.NameStr || hubSession.SecureNATModeBool || hubSession.LinkModeBool {
				continue
			}
			macAddress := ""
			if macT, ok := macTablesMap[hubSession.NameStr]; ok {
				macAddress, _ = pkg.BinToStr(macT.MacAddressBin)
				macAddress = pkg.FormatMacAddress(macAddress)
			}
			startTime, _ := time.Parse("2006-01-02T15:04:05.000Z", hubSession.CreatedTimeDt)
			endTime, _ := time.Parse("2006-01-02T15:04:05.000Z", hubSession.LastCommTimeDt)
			sessions = append(sessions, &response.Session{
				Account:     hubSession.UsernameStr,
				SessionName: hubSession.NameStr,
				ClientIp:    hubSession.ClientIp,
				ClientMac:   macAddress,
				StartTime:   utilTime.TimeFormat(startTime),
				EndTime:     utilTime.TimeFormat(endTime),
			})
		}
	}

	return
}
func (api *Api) HubSessionDelete(hubName string, sessionName string) (err error) {
	err = api.Call(methods.NewDeleteSession(hubName, sessionName), nil)
	return
}
func (api *Api) HubDelete(hubName string) (err error) {
	err = api.Call(methods.NewDeleteHub(hubName), nil)
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
	err = api.Call(m, nil)
	if nil != err {
		return
	}

	if onLine {
		err = api.Call(methods.NewSetLinkOnline(hubName, linkName), nil)
	} else {
		err = api.Call(methods.NewSetLinkOffline(hubName, linkName), nil)
	}

	return
}

func (api *Api) HubLinkDelete(hubName string, linkName string) (err error) {
	err = api.Call(methods.NewDeleteLink(hubName, linkName), nil)

	return
}

func (api *Api) AccountCreate(hubName string, account string, password string) (user *response.ResCreateUser, err error) {

	m := methods.NewCreateUser(hubName, account, account, "", nil, 1)
	m.Params.Auth_Passoword = password

	user = &response.ResCreateUser{}
	err = api.Call(m, user)

	if nil != err {
		err = fmt.Errorf("HUB: %s 创建用户 %s 失败,%+v", hubName, account, err)
	}

	if "" == user.NameStr {
		user = nil
	}
	return
}

func (api *Api) AccountDelete(hubName string, account string) (err error) {

	m := methods.NewDeleteUser(hubName, account)
	err = api.Call(m, nil)
	return
}

func ConvStringPort(str string) (port int) {
	port, err := strconv.Atoi(str)
	if nil != err {
		port = 5555
	}
	return
}
