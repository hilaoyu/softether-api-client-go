package methods

import (
	"bytes"
	"encoding/json"
	"github.com/hilaoyu/softether-api-client-go/pkg"
)

type DeleteUser struct {
	pkg.Base
	Params *DeleteUserParams `json:"params"`
}

func (g *DeleteUser) Parameter() pkg.Params {
	return g.Params
}

func NewDeleteUser(hub, name string) *DeleteUser {
	return &DeleteUser{
		Base:   pkg.NewBase("DeleteUser"),
		Params: newDeleteUserParams(hub, name),
	}
}

func (g *DeleteUser) Name() string {
	return g.Base.Name
}

func (g *DeleteUser) GetId() int {
	return g.Id
}

func (g *DeleteUser) SetId(id int) {
	g.Base.Id = id
}

func (g *DeleteUser) Marshall() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

type DeleteUserParams struct {
	HubName string `json:"HubName_str"`
	Name    string `json:"Name_str"`
}

func newDeleteUserParams(hub, name string) *DeleteUserParams {
	return &DeleteUserParams{
		HubName: hub,
		Name:    name,
	}
}

func (p *DeleteUserParams) Tags() []string {
	tags := []string{
		"HubName_str",
		"Name_str",
	}
	return tags
}
