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

	client, err := indicoserviceauth.NewClient(context)
	if err != nil {
		log.Fatalf(err.Error())
	}

	generateAndValidate(client)
	validateThird(client, "765637")
}

func generateAndValidate(client *indicoserviceauth.Client) {
	response, err := client.GenerateOTP()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Generate OTP Response:")
	fmt.Printf("%+v\n\n", response)

	responsev, err := client.ValidateOTP(response.Token, true)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Validate OTP Response:")
	fmt.Printf("%+v\n\n", responsev)
}

func validateThird(client *indicoserviceauth.Client, token string) {
	responsev, err := client.ValidateOTP(token, false)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Validate Thrid OTP Response:")
	fmt.Printf("%+v\n\n", responsev)

}
