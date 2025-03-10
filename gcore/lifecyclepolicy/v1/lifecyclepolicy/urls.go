package lifecyclepolicy

import (
	"strconv"

	gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"
)

func getURL(c *gcorecloud.ServiceClient, id int) string {
	return c.ServiceURL(strconv.Itoa(id))
}

func listURL(c *gcorecloud.ServiceClient) string {
	return c.ServiceURL()
}

func deleteURL(c *gcorecloud.ServiceClient, id int) string {
	return c.ServiceURL(strconv.Itoa(id))
}

func createURL(c *gcorecloud.ServiceClient) string {
	return c.ServiceURL()
}

func updateURL(c *gcorecloud.ServiceClient, id int) string {
	return c.ServiceURL(strconv.Itoa(id))
}

func addVolumesURL(c *gcorecloud.ServiceClient, id int) string {
	return c.ServiceURL(strconv.Itoa(id), "add_volumes_to_policy")
}

func removeVolumesURL(c *gcorecloud.ServiceClient, id int) string {
	return c.ServiceURL(strconv.Itoa(id), "remove_volumes_from_policy")
}

func addSchedulesURL(c *gcorecloud.ServiceClient, id int) string {
	return c.ServiceURL(strconv.Itoa(id), "add_schedules")
}

func removeSchedulesURL(c *gcorecloud.ServiceClient, id int) string {
	return c.ServiceURL(strconv.Itoa(id), "remove_schedules")
}

func estimateURL(c *gcorecloud.ServiceClient) string {
	return c.ServiceURL("estimate_max_policy_usage")
}
