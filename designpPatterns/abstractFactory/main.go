package main

import (
	ds "stash.experticity.com/go/TestProject_Golang/designpPatterns/abstractFactory/downloadService"
	js "stash.experticity.com/go/TestProject_Golang/designpPatterns/abstractFactory/jobService"

)

func main() {

	dwnSvc := ds.MakeService("mysql")
	//reqStr := "{\"jobType\":\"download_report\"}"
	var req js.AbstractCreateJobRequest = js.AbstractCreateJobRequest{"download_report"}
	var req2 js.AbstractCreateJobRequest = js.AbstractCreateJobRequest{"download_audience"}
	//err := json.Unmarshal([]byte(reqStr), &req)
	//if err != nil{
	//	fmt.Println("Invalid JSON")
	//}
	dwnSvc.Post(req)

	dwnSvc.Post(req2)
}

