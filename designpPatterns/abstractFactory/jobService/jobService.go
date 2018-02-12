package jobService

import (
	"fmt"
)

type JobService interface {
	Post(req ICreateJobRequest)
}

type JobServiceImpl struct {
	DB                 string
}

type ICreateJobRequest interface{
	GetJobType() string
}

type AbstractCreateJobRequest struct {
	JobType string `json:"jobType"`
}

type CreateJobRequest struct {
	AbstractCreateJobRequest
	Email string
	CampaignId string
}

func (a AbstractCreateJobRequest) GetJobType() string {
	return a.JobType
}

func (c CreateJobRequest) GetJobType() string {
	return c.JobType
}

func (m JobServiceImpl) Post(req ICreateJobRequest) {

	fmt.Println("started the job of type: "+req.GetJobType())
	//{
	//
	//} else {
	//	fmt.Println("Invalid data type")
	//}
}