package main

import (
	"stash.experticity.com/go/TestProject_Golang/databases/dynamoDB"
	"fmt"

	"strconv"
)

func main() {

	//create some users
	myUsers  := []dynamoDB.User{}

	//initialize users
	count := 1
	//Note : range copies values into temp variable u and hence any changes made to "u" wouldn't reflect
	//on original collection, hence if you want to alter the values of collection in loop make use of
	//for instead range in golang.

	//for _,u := range myUsers {
	for i := 0 ; i < len(myUsers); i++ {
		u := &myUsers[i]
		u.FirstName = "shrikant"+strconv.Itoa(count)
		u.LastName = "prabhu"+strconv.Itoa(count)
		count = count + 1
		fmt.Println(u)
	}

	//populate user details from dynamoDB
	myUsersPopulated := dynamoDB.GetUserDeviceDataFromDB(myUsers)

	fmt.Println(myUsersPopulated)
}