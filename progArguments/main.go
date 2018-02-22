package main

import (
	"fmt"
	"flag"
	"encoding/json"
	"strconv"
)

type Org struct {
	Id   int    `json:"orgId"`
	Name string `json:"orgName"`
}

func main() {

	//note: pass following variable value as program argument -filters=<filter_fail> or -filters=<filter_success>
	//for some reason flag removes quotes from the string.
	//filters_fail := `{"orgName":"EXPERTICITY","orgId":1}`
	//filters_success := "{\"orgName\":\"EXPERTICITY\",\"orgId\":1}"

	filtersPtr := flag.String("filters", "", "filter criteria")
	flag.Parse()
	fmt.Println("filters: "+*filtersPtr)

	dto := Org{}
	err := json.Unmarshal([]byte(*filtersPtr), &dto)

	if err != nil {
		fmt.Println("Invalid JSON received")
	} else {
		fmt.Println("Valid JSON received")
		fmt.Println("Org Id : " + strconv.Itoa(dto.Id))
		fmt.Println("Org Name : " + dto.Name)
	}
}