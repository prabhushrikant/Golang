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

type ReportCreateJobRequest struct {
	AbstractCreateJobRequest
	Email string
	CampaignId string
}

type AudienceCreateJobRequest struct {
	AbstractCreateJobRequest
	Email string
	FilterData string
}

func (a AbstractCreateJobRequest) GetJobType() string {
	return a.JobType
}

//func (c AudienceCreateJobRequest) GetJobType() string {
//	return c.JobType
//}

func (m JobServiceImpl) Post(req ICreateJobRequest) {

	fmt.Println("started the job of type: "+req.GetJobType())
	//{
	//
	//} else {
	//	fmt.Println("Invalid data type")
	//}
}