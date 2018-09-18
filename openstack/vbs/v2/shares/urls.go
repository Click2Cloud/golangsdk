package shares

import "github.com/huaweicloud/golangsdk"

const shareRootPath = "os-vendor-backup-sharing"

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, shareRootPath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(c.ProjectID, shareRootPath, id)
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, shareRootPath, "detail")
}
