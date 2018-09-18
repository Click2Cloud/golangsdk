package tags

import "github.com/huaweicloud/golangsdk"

const RootPath = "backuppolicy"

func commonURL(c *golangsdk.ServiceClient, policyID string) string {
	return c.ServiceURL(c.ProjectID, RootPath, policyID, "tags")
}

func deleteURL(c *golangsdk.ServiceClient, policyID string, key string) string {
	return c.ServiceURL(c.ProjectID, RootPath, policyID, "tags", key)
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, RootPath, "resource_instances", "action")
}

func actionURL(c *golangsdk.ServiceClient, policyID string) string {
	return c.ServiceURL(c.ProjectID, RootPath, policyID, "tags", "action")
}
