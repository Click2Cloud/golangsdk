package testing

import "github.com/huaweicloud/golangsdk/openstack/cce/v3/nodes"

const Output = `{
    "kind": "Host",
    "apiVersion": "v3",
    "metadata": {
        "name": "c2c-hostname"        
    },
    "spec": {
        "flavor": "s3.large.2",
        "az": "cn-east-2a",
        "login": {
            "sshKey": "c2c-keypair"
        },
        "rootVolume": {
            "volumetype": "SATA",
            "size": 40
        },
        "dataVolumes": [
            {
                "volumetype": "SATA",
                "size": 100
            }
        ]
}
}`

var Expected = &nodes.Nodes{
	Kind:       "Host",
	Apiversion: "v3",
	Metadata:   nodes.Metadata{Name: "c2c-hostname"},
	Spec: nodes.Spec{
		Flavor: "s3.large.2",
		Az:     "cn-east-2a",
		Login: nodes.LoginSpec{
			SshKey: "c2c-keypair",
		},
		RootVolume: nodes.VolumeSpec{
			VolumeType: "SATA",
			Size:       40,
		},
		DataVolumes: []nodes.VolumeSpec{
			{
				VolumeType: "SATA",
				Size:       100,
			},
		},
	},
}
