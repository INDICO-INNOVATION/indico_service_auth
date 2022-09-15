package go_mfaservice

import (
	"os"

	mfaClient "github.com/INDICO-INNOVATION/indico_service_auth/client/mfa"
	grpcHelper "github.com/INDICO-INNOVATION/indico_service_auth/pkg/grpc"
	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/helpers"
)

var mfaservice = mfaClient.NewMFAServiceClient(grpcHelper.Connect(os.Getenv("IAM_SERVER")))

func GenerateOTP(clientID string, clientSecret string) (*mfaClient.GenerateTOTPTokenResponse, error) {
	context, cancel := helpers.InitContext()
	defer cancel()

	otpRequest := &mfaClient.GenerateTOTPTokenRequest{
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}

	return mfaservice.GenerateTOTPToken(context, otpRequest)
}

func ValidateOTP(token string, clientID string, clientSecret string) (*mfaClient.ValidateTOTPTokenResponse, error) {
	context, cancel := helpers.InitContext()
	defer cancel()

	validateRequest := &mfaClient.ValidateTOTPTokenRequest{
		Token:        token,
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}

	return mfaservice.ValidateTOTPToken(context, validateRequest)
}

func GenerateSecretKey(clientID string) (*mfaClient.TOTPSecretResponse, error) {
	context, cancel := helpers.InitContext()
	defer cancel()

	secretRequest := &mfaClient.GenerateTOTPTokenRequest{
		ClientId: clientID,
	}

	return mfaservice.GenerateSecretKey(context, secretRequest)
}
