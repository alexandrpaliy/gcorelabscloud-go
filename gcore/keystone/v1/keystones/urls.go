package keystones

import (
	"strconv"

	gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"
)

func resourceURL(c *gcorecloud.ServiceClient, id int) string {
	return c.BaseServiceURL("keystones", strconv.Itoa(id))
}

func rootURL(c *gcorecloud.ServiceClient) string {
	return c.BaseServiceURL("keystones")
}

func getURL(c *gcorecloud.ServiceClient, id int) string {
	return resourceURL(c, id)
}

func listURL(c *gcorecloud.ServiceClient) string {
	return rootURL(c)
}

func createURL(c *gcorecloud.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *gcorecloud.ServiceClient, id int) string {
	return resourceURL(c, id)
}
