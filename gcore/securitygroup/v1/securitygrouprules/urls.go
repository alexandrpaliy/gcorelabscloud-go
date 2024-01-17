package securitygrouprules

import gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"

func resourceURL(c *gcorecloud.ServiceClient, id string) string {
	return c.ServiceURL(id)
}

func updateURL(c *gcorecloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *gcorecloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
