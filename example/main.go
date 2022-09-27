package main

import (
	"fmt"
	"log"

	indicoserviceauth "github.com/INDICO-INNOVATION/indico_service_auth"
	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/helpers"
)

func main() {
	context, cancel := helpers.InitContext()
	defer cancel()

	client, err := indicoserviceauth.NewClient(context, "mfa.use")
	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.GenerateOTP()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Generate OTP Response:")
	fmt.Printf("%+v\n", response)

	responsev, err := client.ValidateOTP(response.Token)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Validate OTP Response:")
	fmt.Printf("%+v\n", responsev)
}
