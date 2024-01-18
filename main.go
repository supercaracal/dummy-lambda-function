package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	logger = log.New(os.Stderr, "", 0)
)

func main() {
	// https://github.com/aws/aws-lambda-go/tree/main/events
	switch os.Getenv("EVENT") {
	case "Kinesis":
		lambda.Start(Kinesis)
	case "S3":
		lambda.Start(S3)
	default:
		logger.Fatal("EVENT environment variable required")
	}
}

// Kinesis consumes the records of the stream
func Kinesis(ctx context.Context, ev events.KinesisEvent) error {
	for _, record := range ev.Records {
		logger.Print(record.Kinesis.Data)
	}

	return nil
}

// S3 handles the S3 events
func S3(ctx context.Context, ev events.S3Event) error {
	for _, record := range ev.Records {
		logger.Printf("bucket: %s, object: %s", record.S3.Bucket.Name, record.S3.Object.Key)
	}

	return nil
}
