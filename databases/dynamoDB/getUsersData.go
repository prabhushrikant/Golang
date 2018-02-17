package dynamoDB

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"log"
)

func GetUserDeviceDataFromDB(users []User) []User {
	var usersToReturn = users
	var dynamoDBSvc = setupDynamoDB()
	var dynamoDBTable = "user_devices"

	for _, u := range usersToReturn {

		//initialize variables for new user , to achieve do-while behavior
		firstTime := true
		exclusiveStartKey := map[string]*dynamodb.AttributeValue{}

		for firstTime == true || len(exclusiveStartKey) > 0 {
			firstTime = false
			input := &dynamodb.QueryInput{
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
					":uuid1": {
						S: aws.String(u.UserUUID),
					},
				},
				//Because "token" is DynamoDB keyword, can't be used directly in ProjectionExpression instead
				//used with a name placeholder
				ExpressionAttributeNames: map[string]*string{
					"#Proj1": aws.String("token"),
				},
				KeyConditionExpression: aws.String("userUUID = :uuid1"),
				ProjectionExpression:   aws.String("#Proj1,deviceType"),
				TableName:              aws.String(dynamoDBTable),
			}

			if len(exclusiveStartKey) > 0 {
				input.ExclusiveStartKey = exclusiveStartKey
			}

			result, err := dynamoDBSvc.Query(input)
			if err != nil {
				if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					case dynamodb.ErrCodeProvisionedThroughputExceededException:
						log.Fatalln("dynamoDB error:"+dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
					case dynamodb.ErrCodeResourceNotFoundException:
						log.Fatalln("dynamoDB error:"+dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
					case dynamodb.ErrCodeInternalServerError:
						log.Fatalln("dynamoDB error:"+dynamodb.ErrCodeInternalServerError, aerr.Error())
					default:
						log.Fatalln("dynamoDB error:" + aerr.Error())
					}
				} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					log.Fatalln("dynamoDB error:" + err.Error())
				}
			}

			if *result.Count > 0 {

				//fmt.Println("result=" + result.String())
				//populate user device information
				if len(u.Devices) == 0 {
					u.Devices = []DeviceInfo{}
				}

				for _, item := range result.Items {
					var newDevice = DeviceInfo{}
					newDevice.DeviceId = *item["token"].S
					newDevice.DeviceType = *item["deviceType"].S

					u.Devices = append(u.Devices, newDevice)
				}
			}

			if len(result.LastEvaluatedKey) > 0 {

				//get for next page of results for this user
				exclusiveStartKey = result.LastEvaluatedKey
			}
		}

	}

	return usersToReturn
}