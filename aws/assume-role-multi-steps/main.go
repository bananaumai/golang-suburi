package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"io/ioutil"
	"log"
	"os"
)

type Setting struct {
	ARN string
	ExternalID string
}

type Settings struct {
	S1 Setting
	S2 Setting
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

	var ss Settings
	err = json.Unmarshal(data, &ss)
	if err != nil {
		log.Fatalf("failed to unmarshal json - %v", err)
	}

	sess1 := session.Must(session.NewSession())
	creds1 := stscreds.NewCredentials(sess1, ss.S1.ARN, func(p *stscreds.AssumeRoleProvider) {
		if ss.S1.ExternalID != "" {
			p.ExternalID = &ss.S1.ExternalID
		}
	})
	v, err := creds1.Get()
	if err != nil {
		log.Fatalf("failed to get credentials - %v", v)
	}
	log.Printf("creds1 : %v", v)

	sess2 := session.Must(session.NewSession(&aws.Config{
		Credentials: creds1,
	}))
	creds2 := stscreds.NewCredentials(sess2, ss.S2.ARN, func(p *stscreds.AssumeRoleProvider) {
		if ss.S2.ExternalID != "" {
			p.ExternalID = &ss.S2.ExternalID
		}
	})
	v, err = creds2.Get()
	if err != nil {
		log.Fatalf("failed to get credentials - %v", v)
	}
	log.Printf("creds2 : %v", v)
}
