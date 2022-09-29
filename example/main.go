package main

import (
	"context"
	"fmt"
	"log"

	indicoserviceauth "github.com/INDICO-INNOVATION/indico_service_auth"
)

func main() {
	client, ctx, err := indicoserviceauth.NewClient()
	if err != nil {
		log.Fatalf(err.Error())
	}

	generateAndValidate(client, ctx)
	validateThird(client, ctx, "643863")
}

func generateAndValidate(client *indicoserviceauth.Client, ctx context.Context) {
	response, err := client.GenerateOTP(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Generate OTP Response:")
	fmt.Printf("%+v\n\n", response)

	responsev, err := client.ValidateOTP(ctx, response.Token, true)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Validate OTP Response:")
	fmt.Printf("%+v\n\n", responsev)
}

func validateThird(client *indicoserviceauth.Client, ctx context.Context, token string) {
	responsev, err := client.ValidateOTP(ctx, token, false)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Validate Third Party OTP Response:")
	fmt.Printf("%+v\n\n", responsev)

	ctx.Done()
}
