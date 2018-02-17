package utils

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
)

//singleton object
var awsSession *session.Session = nil
var once sync.Once

func createAWSSession() {

	// SETUP: load properties, parse the flags and seed random

	//NOTE - PROVIDE YOUR AWS KEYS HERE.
	os.Setenv("AWS_ACCESS_KEY_ID", "")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "")
	os.Setenv("AWS_REGION", "")

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	awsSession = session.Must(session.NewSession())
	creds, err := awsSession.Config.Credentials.Get()
	if err != nil {
		fmt.Println("aws session credentials denied, sucker", creds)
		log.Fatal(err)
	}
}

func GetAWSSession() *session.Session {
	once.Do(createAWSSession)
	return awsSession
}
