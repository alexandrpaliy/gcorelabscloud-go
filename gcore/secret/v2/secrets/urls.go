package secrets

import gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"

func rootURL(c *gcorecloud.ServiceClient) string {
	return c.ServiceURL()
}

func createURL(c *gcorecloud.ServiceClient) string {
	return rootURL(c)
}
