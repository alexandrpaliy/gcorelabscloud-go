package client

import (
	gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"
	"github.com/alexandrpaliy/gcorelabscloud-go/client/common"

	"github.com/urfave/cli/v2"
)

func NewProjectClientV1(c *cli.Context) (*gcorecloud.ServiceClient, error) {
	return common.BuildClient(c, "projects", "v1")
}
