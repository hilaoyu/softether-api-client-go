package methods

import (
	"bytes"
	"encoding/json"
	"github.com/hilaoyu/softether-api-client-go/pkg"
)

type EnumSession struct {
	pkg.Base
	Params *EnumSessionParams `json:"params"`
}

func (g *EnumSession) Parameter() pkg.Params {
	return g.Params
}

func NewEnumSession(hub string) *EnumSession {
	return &EnumSession{
		Base:   pkg.NewBase("EnumSession"),
		Params: newEnumSessionParams(hub),
	}
}

func (g *EnumSession) Name() string {
	return g.Base.Name
}

func (g *EnumSession) GetId() int {
	return g.Id
}

func (g *EnumSession) SetId(id int) {
	g.Base.Id = id
}

func (g *EnumSession) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type EnumSessionParams struct {
	HubName string `json:"HubName_str"`
}

func newEnumSessionParams(hub string) *EnumSessionParams {
	return &EnumSessionParams{
		HubName: hub,
	}
}
func (p *EnumSessionParams) Tags() []string {
	return []string{"HubName_str"}
}
