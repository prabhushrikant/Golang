package audienceDownloadService

import (
	js "stash.experticity.com/go/TestProject_Golang/designpPatterns/abstractFactory/jobService"
)

//composition (Has-a Relationship than Is-A relationship like in Java)
type AudienceDownloadService interface {
	js.JobService
}

type AudienceDownloadServiceImpl struct {
	//downloadService ds.DownLoadService
	jobService js.JobService
}

func MakeService(DB string) AudienceDownloadService {
	return AudienceDownloadServiceImpl{
			js.JobServiceImpl{
				DB,
			},
	}
}

func (r AudienceDownloadServiceImpl) Post(req js.ICreateJobRequest) {
	r.jobService.Post(req)
}
