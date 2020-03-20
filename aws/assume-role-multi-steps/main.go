package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type AssumeRoleConf struct {
	ARN string
	ExternalID string
}

type Setting struct {
	A1 AssumeRoleConf
	A2 AssumeRoleConf
	TargetStream string
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("setting file should be specified")
	}

	settingFile := args[1]
	data, err := ioutil.ReadFile(settingFile)
	if err != nil {
		log.Fatalf("failed to read setting file from %s - %v", settingFile, err)
	}

	var s Setting
	err = json.Unmarshal(data, &s)
	if err != nil {
		log.Fatalf("failed to unmarshal json - %v", err)
	}

	sess1 := session.Must(session.NewSession())
	creds1 := stscreds.NewCredentials(sess1, s.A1.ARN, func(p *stscreds.AssumeRoleProvider) {
		if s.A1.ExternalID != "" {
			p.ExternalID = &s.A1.ExternalID
		}
	})
	v, err := creds1.Get()
	if err != nil {
		log.Fatalf("failed to get creds1 - %v", v)
	}
	log.Printf("creds1 : %v", v)

	sess2 := session.Must(session.NewSession(&aws.Config{
		Credentials: creds1,
	}))
	creds2 := stscreds.NewCredentials(sess2, s.A2.ARN, func(p *stscreds.AssumeRoleProvider) {
		if s.A2.ExternalID != "" {
			p.ExternalID = &s.A2.ExternalID
		}
	})
	v, err = creds2.Get()
	if err != nil {
		log.Fatalf("failed to get creds2 - %v", v)
	}
	log.Printf("creds2 : %v", v)

	k := kinesis.New(sess2, &aws.Config{
		Credentials: creds2,
	})

	uu, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("failed to generate UUID - %v", err)
	}
	partitionKey := uu.String()

	now := time.Now()

	output, err := k.PutRecord(&kinesis.PutRecordInput{
		StreamName: &s.TargetStream,
		PartitionKey: &partitionKey,
		Data: []byte(now.Format("2006-01-02T15:04:05Z07:00")),
	})
	if err != nil {
		log.Fatalf("failed to put record to kinesis stream - %v", err)
	}

	log.Printf("successfuly put Record - %s", output.String())
}
