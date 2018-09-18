package backups

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/huaweicloud/golangsdk"
)

type JobResponse struct {
	JobID string `json:"job_id"`
}

type JobStatus struct {
	Status     string            `json:"status"`
	Entities   map[string]string `json:"entities"`
	JobID      string            `json:"job_id"`
	JobType    string            `json:"job_type"`
	BeginTime  time.Time         `json:"-"`
	EndTime    time.Time         `json:"-"`
	ErrorCode  string            `json:"error_code"`
	FailReason string            `json:"fail_reason"`
	SubJobs    []JobStatus       `json:"sub_jobs"`
}

type JobResult struct {
	golangsdk.Result
}

func (r JobResult) ExtractJobResponse() (*JobResponse, error) {
	job := new(JobResponse)
	err := r.ExtractInto(job)
	return job, err
}

// UnmarshalJSON overrides the default, to convert the JSON API response into our JobStatus struct
func (r *JobStatus) UnmarshalJSON(b []byte) error {
	type tmp JobStatus
	var s struct {
		tmp
		BeginTime golangsdk.JSONRFC3339Milli `json:"begin_time"`
		EndTime   golangsdk.JSONRFC3339Milli `json:"end_time"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = JobStatus(s.tmp)

	r.BeginTime = time.Time(s.BeginTime)
	r.EndTime = time.Time(s.EndTime)

	return nil
}

func (r JobResult) ExtractJobStatus() (*JobStatus, error) {
	job := new(JobStatus)
	err := r.ExtractInto(job)
	return job, err
}

func WaitForJobSuccess(client *golangsdk.ServiceClient, secs int, jobID string) error {

	client.Endpoint = strings.Replace(client.Endpoint, "v2", "v1", 1)
	return golangsdk.WaitFor(secs, func() (bool, error) {
		job := new(JobStatus)
		_, err := client.Get(client.ServiceURL(client.ProjectID, "jobs", jobID), &job, nil)
		if err != nil {
			return false, err
		}
		fmt.Printf("JobStatus: %+v.\n", job)

		if job.Status == "SUCCESS" {
			return true, nil
		}
		if job.Status == "FAIL" {
			err = fmt.Errorf("Job failed with code %s: %s.\n", job.ErrorCode, job.FailReason)
			return false, err
		}

		return false, nil
	})
}

func GetJobEntity(client *golangsdk.ServiceClient, jobId string, label string) (interface{}, error) {

	client.Endpoint = strings.Replace(client.Endpoint, "v2", "v1", 1)
	job := new(JobStatus)
	_, err := client.Get(client.ServiceURL(client.ProjectID, "jobs", jobId), &job, nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("JobStatus: %+v.\n", job)

	if job.Status == "SUCCESS" {
		if e := job.Entities[label]; e != "" {
			return e, nil
		}
	}

	return nil, fmt.Errorf("Unexpected conversion error in GetJobEntity.")
}
