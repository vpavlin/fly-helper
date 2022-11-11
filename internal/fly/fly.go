package fly

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
	"github.com/superfly/flyctl/client"
	"github.com/superfly/flyctl/flyctl"
)

func NewFly() *client.Client {
	flyctl.InitConfig()
	return client.New()
}

type FlyToml struct {
	App string `toml:"app"`
}

func NewFlyToml(path string) (*FlyToml, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var fly FlyToml
	err = toml.Unmarshal(data, &fly)
	if err != nil {
		return nil, err
	}

	return &fly, nil
}
