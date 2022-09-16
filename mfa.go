package go_mfaservice

import (
	mfaClient "github.com/INDICO-INNOVATION/indico_service_auth/client/mfa"
	grpcHelper "github.com/INDICO-INNOVATION/indico_service_auth/pkg/grpc"
	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/helpers"
)

var mfaservice = mfaClient.NewMFAServiceClient(grpcHelper.Connect())

func GenerateOTP(clientID string, clientSecret string) (*mfaClient.GenerateOTPTokenResponse, error) {
	context, cancel := helpers.InitContext()
	defer cancel()

	otpRequest := &mfaClient.GenerateOTPTokenRequest{
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}

	return mfaservice.GenerateOTPToken(context, otpRequest)
}

func ValidateOTP(token string, clientID string, clientSecret string) (*mfaClient.ValidateOTPTokenResponse, error) {
	context, cancel := helpers.InitContext()
	defer cancel()

	validateRequest := &mfaClient.ValidateOTPTokenRequest{
		Token:        token,
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}

	return mfaservice.ValidateOTPToken(context, validateRequest)
}

func GenerateSecretKey(clientID string) (*mfaClient.OTPSecretResponse, error) {
	context, cancel := helpers.InitContext()
	defer cancel()

	secretRequest := &mfaClient.GenerateOTPTokenRequest{
		ClientId: clientID,
	}

	return mfaservice.GenerateSecretKey(context, secretRequest)
}
