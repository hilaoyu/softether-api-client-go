package methods

import (
	"bytes"
	"encoding/json"
	"github.com/hilaoyu/softether-api-client-go/pkg"
)

type DeleteSession struct {
	pkg.Base
	Params *DeleteSessionParams `json:"params"`
}

func (g *DeleteSession) Parameter() pkg.Params {
	return g.Params
}

func NewDeleteSession(hub, name string) *DeleteSession {
	return &DeleteSession{
		Base:   pkg.NewBase("DeleteSession"),
		Params: newDeleteSessionParams(hub, name),
	}
}

func (g *DeleteSession) Name() string {
	return g.Base.Name
}

func (g *DeleteSession) GetId() int {
	return g.Id
}

func (g *DeleteSession) SetId(id int) {
	g.Base.Id = id
}

func (g *DeleteSession) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type DeleteSessionParams struct {
	HubName string `json:"HubName_str"`
	Name    string `json:"Name_str"`
}

func newDeleteSessionParams(hub, name string) *DeleteSessionParams {
	return &DeleteSessionParams{
		HubName: hub,
		Name:    name,
	}
}

func (p *DeleteSessionParams) Tags() []string {
	tags := []string{
		"HubName_str",
		"Name_str",
	}
	return tags
}
