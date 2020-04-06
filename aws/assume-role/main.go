package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

var (
	arn string
	externalID string
)

func init() {
	flag.StringVar(&arn, "arn", "", "Target ARN")
	flag.StringVar(&externalID, "external_id", "", "External ID")

	flag.Parse()
}

func main() {
	sess := session.Must(session.NewSession())
	creds := stscreds.NewCredentials(sess, arn, func(p *stscreds.AssumeRoleProvider) {
		if externalID != "" {
			p.ExternalID = &externalID
		}
	})

	expiresAt, err := creds.ExpiresAt()
	if err != nil {
		log.Fatalf("failed to get expires at - %v", err)
	}
	fmt.Printf("creds expires at %s\n", expiresAt.Format("2006-01-02T15:04:05Z07:00"))


	v, err := creds.Get()
	if err != nil {
		log.Fatalf("failed to get credentials - %v", err)
	}

	fmt.Printf("creds value %v", v)
}
