package methods

import (
	"bytes"
	"encoding/json"
	"github.com/hilaoyu/softether-api-client-go/pkg"
)

type EnumMacTable struct {
	pkg.Base
	Params *EnumMacTableParams `json:"params"`
}

func NewEnumMacTable(name string) *EnumMacTable {
	return &EnumMacTable{
		Base:   pkg.NewBase("EnumMacTable"),
		Params: newEnumMacTableParams(name),
	}
}
func (m *EnumMacTable) Name() string {
	return m.Base.Name
}

func (m *EnumMacTable) GetId() int {
	return m.Id
}

func (m *EnumMacTable) SetId(id int) {
	m.Base.Id = id
}

func (m *EnumMacTable) Parameter() pkg.Params {
	return m.Params
}

func (m *EnumMacTable) Marshall() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type EnumMacTableParams struct {
	HubName string `json:"HubName_str"`
}

func newEnumMacTableParams(hub string) *EnumMacTableParams {
	return &EnumMacTableParams{HubName: hub}
}

func (p *EnumMacTableParams) Tags() []string {
	return []string{
		"HubName_str",
	}
}
