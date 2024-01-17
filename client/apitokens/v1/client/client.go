package client

import (
	gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"
	"github.com/alexandrpaliy/gcorelabscloud-go/client/common"
	"github.com/alexandrpaliy/gcorelabscloud-go/gcore"
	"github.com/urfave/cli/v2"
)

func NewAPITokenClient(c *cli.Context) (*gcorecloud.ServiceClient, error) {
	// todo refactor it, now apitokens could be generated only with platform client type
	settings, err := gcore.NewGCloudPlatformAPISettingsFromEnv()
	if err != nil {
		return nil, err
	}

	ao, err := gcore.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	ao.APIURL = settings.AuthURL
	return common.BuildAPITokenClient(ao)
}
