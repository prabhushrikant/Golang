package reportDownloadService

import (
	js "stash.experticity.com/go/TestProject_Golang/designpPatterns/abstractFactory/jobService"
)

//composition (Has-a Relationship than Is-A relationship like in Java)
type ReportDownloadService interface {
	js.JobService
}

type ReportDownloadServiceImpl struct {
	//downloadService ds.DownLoadService
	jobService js.JobService
}

func MakeService(DB string) ReportDownloadService {
	return ReportDownloadServiceImpl{
			js.JobServiceImpl{
				DB,
			},
	}
}

func (r ReportDownloadServiceImpl) Post(req js.ICreateJobRequest) {
	r.jobService.Post(req)
}
