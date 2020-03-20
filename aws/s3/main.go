package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "file")

	var bucket string
	flag.StringVar(&bucket, "b", "", "bucket")

	var key string
	flag.StringVar(&key, "k", "", "key")

	flag.Parse()

	sess := session.Must(session.NewSession())

	uploader := s3manager.NewUploader(sess)

	f, err  := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open file %q, %v", filename, err)
		os.Exit(1)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   f,
	})

	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
		os.Exit(1)
	}

	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
}
