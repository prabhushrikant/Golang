package audienceDownloadService

import (
	js "stash.experticity.com/go/TestProject_Golang/designpPatterns/abstractFactory/jobService"
	"fmt"
	"errors"
)

type AudienceDownloadService interface {
	js.JobService
	convertCreateRequest(req js.ICreateJobRequest) (*js.AudienceCreateJobRequest, error)
}

//composition (Has-a Relationship than Is-A relationship like in Java)
type AudienceDownloadServiceImpl struct {
	jobServiceImpl js.JobServiceAudienceImpl
}

func MakeService(DB string) AudienceDownloadService {
	return AudienceDownloadServiceImpl{
			js.JobServiceAudienceImpl{
				js.JobServiceImpl{
				DB,
			},
		},
	}
}


func (r AudienceDownloadServiceImpl) Post(req js.ICreateJobRequest) {
	newReq, err := r.convertCreateRequest(req)
	if err != nil {
		fmt.Println("err - ",err)
	}
	r.jobServiceImpl.Post(*newReq)
}

func (s AudienceDownloadServiceImpl) convertCreateRequest(req js.ICreateJobRequest) (*js.AudienceCreateJobRequest, error) {
	//type assertion
	newReq, ok := req.(js.AudienceCreateJobRequest)
	if ok != true {
		return nil, errors.New("Invalid type specified")
	}
	return &newReq, nil
}