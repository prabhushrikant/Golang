package main

import (
	ds "stash.experticity.com/go/TestProject_Golang/designpPatterns/abstractFactory/downloadService"
	js "stash.experticity.com/go/TestProject_Golang/designpPatterns/abstractFactory/jobService"
)

func main() {

	dwnSvc := ds.MakeService("mysql")
	//reqStr := "{\"jobType\":\"download_report\"}"
	var req js.ReportCreateJobRequest = js.ReportCreateJobRequest{
		js.AbstractCreateJobRequest{"download_report"},
		"prabhushrikant@gmail.com",
		"1234",
	}

	var req2 js.AudienceCreateJobRequest = js.AudienceCreateJobRequest{
		js.AbstractCreateJobRequest{"download_audience"},
		"prabhushrikant@gmail.com",
		"{\"category_id\":{\"val\":[1,2,32],\"op\":\"or\"}}",
	}

	//err := json.Unmarshal([]byte(reqStr), &req)
	//if err != nil{
	//	fmt.Println("Invalid JSON")
	//}
	dwnSvc.Post(req)

	dwnSvc.Post(req2)
}

