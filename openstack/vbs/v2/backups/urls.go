package backups

import "github.com/huaweicloud/golangsdk"

const (
	cloudNativeRootPath = "cloudbackups"
	osNativeRootPath    = "backups"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, cloudNativeRootPath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(c.ProjectID, osNativeRootPath, id)
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, osNativeRootPath, "detail")
}

func restoreURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(c.ProjectID, osNativeRootPath, id, "restore")
}
