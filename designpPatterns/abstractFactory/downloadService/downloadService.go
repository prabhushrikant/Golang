package downloadService

import (
	"fmt"
	js "stash.experticity.com/go/TestProject_Golang/designpPatterns/abstractFactory/jobService"
	"stash.experticity.com/go/TestProject_Golang/designpPatterns/abstractFactory/reportDownloadService"
	"stash.experticity.com/go/TestProject_Golang/designpPatterns/abstractFactory/audienceDownloadService"
)

type DownloadService struct {
	//jobService js.JobServiceImpl
	DB string
}

func MakeService(DB string) DownloadService {
	return DownloadService{
			DB,
	}
}

func (d DownloadService) downloadServiceFactory(req js.ICreateJobRequest) (svc js.JobService, err string) {

	if req.GetJobType() == "download_report"{
		return reportDownloadService.MakeService(d.DB), ""
	} else if req.GetJobType() == "download_audience"{
		return audienceDownloadService.MakeService(d.DB), ""
	} else {
		return nil, "invalid job type received"
	}
}

func (d DownloadService) Post(req js.ICreateJobRequest) {

	svc,err := d.downloadServiceFactory(req)

	if err != "" {
		fmt.Println("Some error in factory : ",err)
	}

	svc.Post(req)
}




