package fly

import (
	"github.com/superfly/flyctl/client"
	"github.com/superfly/flyctl/flyctl"
)

func NewFly() *client.Client {
	flyctl.InitConfig()
	return client.New()
}
