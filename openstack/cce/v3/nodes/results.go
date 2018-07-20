package nodes

import (
	"github.com/huaweicloud/golangsdk"
)

// Describes the Node Structure of cluster
type ListNodeResponse struct {
	// API type, fixed value "List"
	Kind       string  `json:"kind"`
	// API version, fixed value "v3"
	Apiversion string  `json:"apiVersion"`
	// all Clusters
	Nodes      []Nodes `json:"items"`
}

// Individual nodes of the cluster
type Nodes struct {
	//  API type, fixed value " Host "
	Kind       string   `json:"kind"`
	// API version, fixed value v3
	Apiversion string   `json:"apiVersion"`
	// Node metadata
	Metadata   Metadata `json:"metadata"`
	// Node detailed parameters
	Spec       Spec     `json:"spec"`
	// Node status information
	Status     Status   `json:"status"`
}

// Metadata required to create a node
type Metadata struct {
	// Node name
	Name string `json:"name"`
	// Node ID
	Uid string `json:"uid"`
	// Node tag, key value pair format
	Labels map[string]string `json:"labels,omitempty"`
	// Node annotation, keyvalue pair format
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Describes Nodes specification
type Spec struct {
	// Node specifications
	Flavor      string       `json:"flavor" required:"true"`
	// The value of the available partition name
	Az          string       `json:"az" required:"true"`
	// Node login parameters
	Login       LoginSpec    `json:"login" required:"true"`
	// System disk parameter of the node
	RootVolume  VolumeSpec   `json:"rootVolume" required:"true"`
	// The data disk parameter of the node must currently be a disk
	DataVolumes []VolumeSpec `json:"dataVolumes" required:"true"`
	// Elastic IP parameters of the node
	PublicIP    PublicIPSpec `json:"publicIP,omitempty"`
	// The billing mode of the node: the value is 0 (on demand)
	BillingMode int          `json:"billingMode,omitempty"`
	// Number of nodes when creating in batch
	Count       int          `json:"count" required:"true"`
	// Extended parameter
	ExtendParam string       `json:"extendParam,omitempty"`
}

// Gives the current status of the node
type Status struct {
	// The state of the Node
	Phase     string `json:"phase"`
	// The virtual machine ID of the node in the ECS
	ServerID  string `json:"ServerID"`
	// Elastic IP of the node
	PublicIP  string `json:"PublicIP"`
	//Private IP of the node
	PrivateIP string `json:"privateIP"`
	// The ID of the Job that is operating asynchronously in the Node
	JobID string `json:"jobID"`
	// Reasons for the Node to become current
	Reason  string `json:"reason"`
	// Details of the node transitioning to the current state
	Message string `json:"message"`
	// The status of each component in the Node
	Conditions Conditions `json:"conditions"`
}

type LoginSpec struct {
	// Select the key pair name when logging in by key pair mode
	SshKey string `json:"sshKey" required:"true"`
}

type VolumeSpec struct {
	// Disk size in GB
	Size        int    `json:"size" required:"true"`
	// Disk type
	VolumeType  string `json:"volumetype" required:"true"`
	// Disk extension parameter
	ExtendParam string `json:"extendParam,omitempty"`
}

type PublicIPSpec struct {
	// List of existing elastic IP IDs
	Ids   []string `json:"ids,omitempty"`
	// The number of elastic IPs to be dynamically created
	Count int      `json:"count,omitempty"`
	// Elastic IP parameters
	Eip   EipSpec  `json:"eip,omitempty"`
}

type EipSpec struct {
	// The value of the iptype keyword
	IpType    string        `json:"iptype" required:"true"`
	// Elastic IP bandwidth parameters
	Bandwidth BandwidthOpts `json:"bandwidth" required:"true"`
}

type BandwidthOpts struct {
	ChargeMode string `json:"chargemode,omitempty"`
	Size       int    `json:"size" required:"true"`
	ShareType  string `json:"sharetype" required:"true"`
}

type Conditions struct {
	// The type of component
	Type string `json:"type"`
	// The state of the component
	Status string `json:"status"`
	// The reason that the component becomes current
	Reason string `json:"reason"`
}

// Describes the Job Structure
type Job struct {
	Kind       string      `json:"kind"`
	Apiversion string      `json:"apiVersion"`
	Metadata   JobMetadata `json:"metadata"`
	Spec       JobSpec     `json:"spec"`
}

type JobMetadata struct {
	Uid string `json:"uid"`
}

type JobSpec struct {
	Type         string    `json:"type"`
	ClusterUID   string    `json:"clusterUID"`
	ResourceName string    `json:"resourceName"`
	SubJobs      []SubJobs `json:"subJobs"`
}

type SubJobs struct {
	Spec SpecJobCluster `json:"spec"`
}

type SpecJobCluster struct {
	SubJobsCluster []SubJobsCluster `json:"subJobs"`
}

type SubJobsCluster struct {
	SpecJobNode SpecJobNode `json:"spec"`
}

type SpecJobNode struct {
	Type       string `json:"type"`
	ClusterUID string `json:"clusterUID"`
	ResourceID string `json:"resourceID"`
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

// ExtractJob is a function that accepts a result and extracts a job.
func (r commonResult) ExtractJob() (*Job, error) {
	var s Job
	err := r.ExtractInto(&s)
	return &s, err
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
