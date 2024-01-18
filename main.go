package main

import (
	"context"
	"log"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	logger = log.New(os.Stderr, "", 0)
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, v interface{}) error {
	// https://github.com/aws/aws-lambda-go/tree/main/events
	switch ev := v.(type) {
	case events.KinesisEvent:
		for _, record := range ev.Records {
			logger.Print(record.Kinesis.Data)
		}
	case events.S3Event:
		for _, record := range ev.Records {
			key, err := url.QueryUnescape(record.S3.Object.Key)
			if err != nil {
				logger.Print(err)
				continue
			}
			logger.Printf("bucket: %s, object: %s", record.S3.Bucket.Name, key)
		}
	default:
		logger.Print("not implemented yet")
	}

	return nil
}
