package aiimages

import (
	gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"
)

func listAIImagesURL(c *gcorecloud.ServiceClient) string {
	return c.ServiceURL()
}
