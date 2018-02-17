package dynamoDB

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"stash.experticity.com/go/TestProject_Golang/utils"
)

func setupDynamoDB() *dynamodb.DynamoDB {
	// Create a DynamoDB client from just a session.
	session := utils.GetAWSSession()
	svc := dynamodb.New(session)
	return svc
}
