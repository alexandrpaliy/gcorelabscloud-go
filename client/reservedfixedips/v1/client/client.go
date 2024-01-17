package client

import (
	"github.com/urfave/cli/v2"

	gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"
	"github.com/alexandrpaliy/gcorelabscloud-go/client/common"
)

func NewReservedFixedIPClientV1(c *cli.Context) (*gcorecloud.ServiceClient, error) {
	return common.BuildClient(c, "reserved_fixed_ips", "v1")
}
