package keypairs

import gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"

func resourceURL(c *gcorecloud.ServiceClient, id string) string {
	return c.BaseServiceURL("keypairs", id)
}

func rootURL(c *gcorecloud.ServiceClient) string {
	return c.BaseServiceURL("keypairs")
}

func getURL(c *gcorecloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listURL(c *gcorecloud.ServiceClient) string {
	return rootURL(c)
}

func createURL(c *gcorecloud.ServiceClient) string {
	return rootURL(c)
}

func deleteURL(c *gcorecloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
