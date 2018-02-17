package jobService

import "fmt"

type JobServiceAudienceImpl struct {
	JobServiceImpl
}

func (a JobServiceAudienceImpl) Post(req AudienceCreateJobRequest) {

	fmt.Println("started the job of type: "+req.GetJobType() + " with filter "+req.FilterData)
	a.insertIntoDatabase(req)
}