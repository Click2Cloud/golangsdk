package testing

const NodeGetResponse = `
{
    "kind": "host",
    "apiVersion": "v1",
    "metadata": {
        "name": "test-cluster-node-1",
        "uuid": "5e6e7641-c288-40a2-a5af-014cc600d838",
        "spaceuuid": "c0b1f563-400a-dec4-9d6f-d69691261a5b",
        "createAt": "2018-05-11 10:04:54.712096939 +0000 UTC",
        "updateAt": "2018-05-11 10:09:59"
    },
    "spec": {
        "clusteruuid": "00ebf544-4db8-4f0c-9359-75115861a63a",
        "clustername": "c2c-test-cluster",
        "privateip": "192.168.0.221",
        "publicip": "80.158.23.56",
        "flavor": "s1.medium",
        "cpu": 1,
        "memory": 4096,
        "hostid": "5e6e7641-c288-40a2-a5af-014cc600d838",
        "az": "eu-de-02",
        "volume": [
            {
                "diskType": "root",
                "diskSize": 40,
                "volumeType": "SATA"
            },
            {
                "diskType": "data",
                "diskSize": 100,
                "volumeType": "SATA"
            }
        ],
        "sshkey": "test-keypair1",
        "status": {
            "capacity": {
                "cpu": "1",
                "memory": "3844320Ki",
                "pods": "110"
            },
            "allocatable": {
                "cpu": "0.452",
                "memory": "3317984Ki",
                "pods": "110"
            },
            "conditions": [
                {
                    "type": "OutOfDisk",
                    "status": "False",
                    "lastHeartbeatTime": "2018-05-11T10:29:58Z",
                    "lastTransitionTime": "2018-05-11T10:17:27Z",
                    "reason": "KubeletHasSufficientDisk",
                    "message": "kubelet has sufficient disk space available"
                },
                {
                    "type": "MemoryPressure",
                    "status": "False",
                    "lastHeartbeatTime": "2018-05-11T10:29:58Z",
                    "lastTransitionTime": "2018-05-11T10:17:27Z",
                    "reason": "KubeletHasSufficientMemory",
                    "message": "kubelet has sufficient memory available"
                },
                {
                    "type": "DiskPressure",
                    "status": "False",
                    "lastHeartbeatTime": "2018-05-11T10:29:58Z",
                    "lastTransitionTime": "2018-05-11T10:17:27Z",
                    "reason": "KubeletHasNoDiskPressure",
                    "message": "kubelet has no disk pressure"
                },
                {
                    "type": "Ready",
                    "status": "True",
                    "lastHeartbeatTime": "2018-05-11T10:29:58Z",
                    "lastTransitionTime": "2018-05-11T10:17:27Z",
                    "reason": "KubeletReady",
                    "message": "kubelet is posting ready status"
                }
            ],
            "addresses": [
                {
                    "type": "InternalIP",
                    "address": "192.168.0.221"
                },
                {
                    "type": "Hostname",
                    "address": "192.168.0.221"
                }
            ],
            "daemonEndpoints": {
                "kubeletEndpoint": {
                    "Port": 10250
                }
            },
            "nodeInfo": {
                "machineID": "f21bcad2f362498d924ac24efcdd3e2a",
                "systemUUID": "CD638027-C922-4AED-A865-6E7D0B98EFB4",
                "bootID": "a74dbac1-df46-4989-9e5c-8093000aeeeb",
                "kernelVersion": "3.10.0-229.42.1.97.x86_64",
                "osImage": "EulerOS V2.0SP1",
                "containerRuntimeVersion": "docker://1.11.2",
                "kubeletVersion": "cce-1.7.3.17+8857de1a267f8f-dirty",
                "kubeProxyVersion": "cce-1.7.3.17+8857de1a267f8f-dirty"
            },
            "images": [
                {
                    "names": [
                        "fluentd:1.11"
                    ],
                    "sizeBytes": 617825114
                },
                {
                    "names": [
                        "ops_agent:node"
                    ],
                    "sizeBytes": 200491883
                },
                {
                    "names": [
                        "kube-proxy:cce-1.7.3.17_8857de1a267f8f-dirty"
                    ],
                    "sizeBytes": 114793936
                },
                {
                    "names": [
                        "gcr.io/google_containers/heapster:v1.4.0"
                    ],
                    "sizeBytes": 73395475
                },
                {
                    "names": [
                        "kubedns-amd64:1.5"
                    ],
                    "sizeBytes": 50815316
                },
                {
                    "names": [
                        "gcr.io/google_containers/addon-resizer:1.7"
                    ],
                    "sizeBytes": 38983736
                },
                {
                    "names": [
                        "quay.io/coreos/etcd:v3.2"
                    ],
                    "sizeBytes": 35749746
                },
                {
                    "names": [
                        "exechealthz-amd64:1.0"
                    ],
                    "sizeBytes": 7115733
                },
                {
                    "names": [
                        "kube-dnsmasq-amd64:1.3"
                    ],
                    "sizeBytes": 5125973
                },
                {
                    "names": [
                        "pause:2.0"
                    ],
                    "sizeBytes": 350164
                }
            ]
        }
    },
    "replicas": 1,
    "status": "ACTIVE",
    "message": "test-cluster-node-1"
}
`
