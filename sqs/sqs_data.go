package sqs

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

//Sqs Data calls AWS SDK
func SqsData() bool {

	now := time.Now()
	// .Format("15:04:05")

	// Create a client
	client := sqs.New(session.New(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	input := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String("https://sqs.us-east-1.amazonaws.com/261472364019/exzeo-print-inbox-file-received-demo"),
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(1),
		// AttributeNames:      aws.([]string(locationNameList:"All" type:"list" flattened:"true"))
	}

	//Using Client to Call API for a message
	resp, err := client.ReceiveMessage(input)
	if err != nil {
		panic(err)
	}
	//Parsing Date to pull just the date
	jsonStr := *resp.Messages[0].Body
	trim := strings.Split(jsonStr, ",")
	event := strings.TrimLeft(trim[3], `"eventTime:"`)
	eventStage1 := strings.TrimRight(event, `"`)
	eventStage2 := strings.Split(eventStage1, "T")

	timeOnlyStage := eventStage2[1]
	timeOnlyStage2 := strings.Split(timeOnlyStage, ".")
	timeOnlyStage3 := strings.Split(timeOnlyStage2[0], ":")
	timeOnly := timeOnlyStage3[0] + ":" + timeOnlyStage3[1]

	dateOnly := eventStage2[0]

	fmt.Println("DateOnly:", dateOnly)
	fmt.Println("TimeOnly:", timeOnly)

	myDate, _ := time.Parse("2006-01-02 15:04", dateOnly+" "+timeOnly)

	timeDiff := myDate.Sub(now)

	fmt.Println("TimeDiff:", timeDiff)

	t := int64(timeDiff / 1000000000)

	if t <= -360000 {
		return true
	}
	return false

}
