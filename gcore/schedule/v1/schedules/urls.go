package schedules

import gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"

func resourceURL(c *gcorecloud.ServiceClient, id string) string {
	return c.ServiceURL(id)
}
