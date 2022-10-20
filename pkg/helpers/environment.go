package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type InnovationCredentials struct {
	Type         string `json:"type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	PrivateKey   string `json:"private_key"`
	ClientEmail  string `json:"client_email"`
	AuthURI      string `json:"auth_uri"`
}

func ParseEnvironment() (*InnovationCredentials, string) {
	var innovationCredentials *InnovationCredentials

	if os.Getenv("INNOVATION_CREDENTIALS") != "" {
		filebyte, err := ioutil.ReadFile(os.Getenv("INNOVATION_CREDENTIALS"))
		if err != nil {
			log.Fatalf("could not find INNOVATION_CREDENTIALS file: %s", err.Error())
		}

		if err = json.Unmarshal([]byte(filebyte), &innovationCredentials); err != nil {
			log.Fatalf("could not parse INNOVATION_CREDENTIALS environment: %s", err.Error())
		}
	} else {
		log.Println("environment variable INNOVATION_CREDENTIALS not set")
	}

	// TODO: Change uri when it will be available in prod
	const AUTH_SERVER = "localhost:50051"

	var authServer string = AUTH_SERVER
	if os.Getenv("AUTH_SERVER") != "" {
		authServer = os.Getenv("AUTH_SERVER")
	}

	return innovationCredentials, authServer
}
