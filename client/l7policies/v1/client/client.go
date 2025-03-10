package client

import (
	gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"
	"github.com/alexandrpaliy/gcorelabscloud-go/client/common"

	"github.com/urfave/cli/v2"
)

func NewL7PoliciesClientV1(c *cli.Context) (*gcorecloud.ServiceClient, error) {
	return common.BuildClient(c, "l7policies", "v1")
}

func NewL7RulesClientV1(c *cli.Context) (*gcorecloud.ServiceClient, error) {
	return NewL7PoliciesClientV1(c)
}
