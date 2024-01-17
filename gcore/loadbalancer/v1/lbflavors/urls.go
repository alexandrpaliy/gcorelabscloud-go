package lbflavors

import gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"

func listURL(c *gcorecloud.ServiceClient) string {
	return c.ServiceURL()
}
