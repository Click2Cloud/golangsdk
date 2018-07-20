package nodes

import (
	"github.com/huaweicloud/golangsdk"
)

//Describes the Node Structure of cluster
type ListNodeResponse struct {
	Kind       string  `json:"kind"`
	Apiversion string  `json:"apiVersion"`
	Nodes      []Nodes `json:"items"`
}

//Individual nodes of the cluster
type Nodes struct {
	Kind       string   `json:"kind"`
	Apiversion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
	Status     Status   `json:"status"`
}

//Metadata required to create a node
type Metadata struct {
	//Node name
	Name string `json:"name"`
	//Node ID
	Uid string `json:"uid"`
	// Node tag, key value pair format
	Labels map[string]string `json:"labels,omitempty"`
	//Node annotation, keyvalue pair format
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Describes Nodes specification
type Spec struct {
	Flavor      string       `json:"flavor" required:"true"`
	Az          string       `json:"az" required:"true"`
	Login       LoginSpec    `json:"login" required:"true"`
	RootVolume  VolumeSpec   `json:"rootVolume" required:"true"`
	DataVolumes []VolumeSpec `json:"dataVolumes" required:"true"`
	PublicIP    PublicIPSpec `json:"publicIP,omitempty"`
	BillingMode int          `json:"billingMode,omitempty"`
	Count       int          `json:"count" required:"true"`
	ExtendParam string       `json:"extendParam,omitempty"`
}

//Gives the current status of the node
type Status struct {
	//The state of the Node
	Phase     string `json:"phase"`
	ServerID  string `json:"ServerID"`
	PublicIP  string `json:"PublicIP"`
	PrivateIP string `json:"privateIP"`
	//The ID of the Job that is operating asynchronously in the Node
	JobID string `json:"jobID"`
	//Reasons for the Node to become current
	Reason  string `json:"reason"`
	Message string `json:"message"`
	//The status of each component in the Node
	Conditions Conditions `json:"conditions"`
}

type LoginSpec struct {
	SshKey string `json:"sshKey" required:"true"`
}

type VolumeSpec struct {
	Size        int    `json:"size" required:"true"`
	VolumeType  string `json:"volumetype" required:"true"`
	ExtendParam string `json:"extendParam,omitempty"`
}

type PublicIPSpec struct {
	Ids   []string `json:"ids,omitempty"`
	Count int      `json:"count,omitempty"`
	Eip   EipSpec  `json:"eip,omitempty"`
}

type EipSpec struct {
	IpType    string        `json:"iptype" required:"true"`
	Bandwidth BandwidthOpts `json:"bandwidth" required:"true"`
}

type BandwidthOpts struct {
	ChargeMode string `json:"chargemode,omitempty"`
	Size       int    `json:"size" required:"true"`
	ShareType  string `json:"sharetype" required:"true"`
}

type Conditions struct {
	//The type of component
	Type string `json:"type"`
	//The state of the component
	Status string `json:"status"`
	//The reason that the component becomes current
	Reason string `json:"reason"`
}

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a node.
func (r commonResult) Extract() (*Nodes, error) {
	var s Nodes
	err := r.ExtractInto(&s)
	return &s, err
}

// ExtractNode is a function that accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func (r commonResult) ExtractNode(opts ListOpts) ([]Nodes, error) {
	var s ListNodeResponse
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return FilterNodes(s.Nodes, opts)
}

func (r commonResult) ExtractJob() (*Job, error) {
	var s Job
	err := r.ExtractInto(&s)
	return &s, err
}

type Job struct {
	Kind       string      `json:"kind"`
	Apiversion string      `json:"apiVersion"`
	Metadata   JobMetadata `json:"metadata"`
	Spec       JobSpec     `json:"spec"`
	Status     JobStatus   `json:"status"`
}

type JobMetadata struct {
	Uid string `json:"uid"`
}

type JobSpec struct {
	Type         string `json:"type"`
	ClusterUID   string `json:"clusterUID"`
	ResourceID   string `json:"resourceID"`
	ResourceName string `json:"resourceName"`
	SubJobs      []Job  `json:"subJobs"`
	OwnerJob     string `json:"ownerJob"`
}

type JobStatus struct {
	Phase   string `json:"phase"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

// ListResult represents the result of a list operation. Call its ExtractNode
// method to interpret it as a Nodes.
type ListResult struct {
	commonResult
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Node.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Node.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Node.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}
