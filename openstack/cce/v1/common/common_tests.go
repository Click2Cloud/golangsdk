package common

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "api/"+ "v1/"
	return sc
}

/*func NodeServiceClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "api/"+ "v1/"
	return sc
}*/