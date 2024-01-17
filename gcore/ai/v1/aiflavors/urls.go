package aiflavors

import (
	gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"
)

func listAIFlavorsURL(c *gcorecloud.ServiceClient) string {
	return c.ServiceURL()
}
