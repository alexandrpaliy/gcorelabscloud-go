package testing

import (
	gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"
	"github.com/alexandrpaliy/gcorelabscloud-go/testhelper"
)

func createClient() *gcorecloud.ServiceClient {
	return &gcorecloud.ServiceClient{
		ProviderClient: &gcorecloud.ProviderClient{AccessTokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
