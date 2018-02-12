package audienceDownloadService

import (
	js "stash.experticity.com/go/TestProject_Golang/designpPatterns/abstractFactory/jobService"
	"fmt"
)

type AudienceDownloadService interface {
	js.JobService
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
	newReq, ok := req.(js.AudienceCreateJobRequest)
	if ok != true {
		fmt.Println("Type conversion error")
	}
	r.jobServiceImpl.Post(newReq)
}
