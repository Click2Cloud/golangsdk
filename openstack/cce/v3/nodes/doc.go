/*
Package nodes enables management and retrieval of nodes
CCE service.

Example to List nodes

	listNodes := nodes.ListOpts{}
	clusterID := "cec124c2-58f1-11e8-ad73-0255ac101926"
	allNodes, err := nodes.List(client,clusterID ).ExtractNode(listNodes)

	if err != nil {
		panic(err)
	}

	for _, node := range allNodes {
		fmt.Printf("%+v\n", node)
	}


*/

package nodes
