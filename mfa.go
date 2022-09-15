package go_mfaservice

import (
	mfaClient "github.com/INDICO-INNOVATION/indico_service_auth/client/mfa"
	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/helpers"
)

type MfaService interface {
	GenerateOTP(clientID string, clientSecret string) (*mfaClient.GenerateTOTPTokenResponse, error)
	ValidateOTP(token string, clientID string, clientSecret string) (*mfaClient.ValidateTOTPTokenResponse, error)
	GenerateSecretKey(clientID string) (*mfaClient.TOTPSecretResponse, error)
}

type mfaservice struct {
	service mfaClient.MFAServiceClient
}

func (mfa *mfaservice) GenerateOTP(clientID string, clientSecret string) (*mfaClient.GenerateTOTPTokenResponse, error) {
	context, cancel := helpers.InitContext()
	defer cancel()

	otpRequest := &mfaClient.GenerateTOTPTokenRequest{
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}

	return mfa.service.GenerateTOTPToken(context, otpRequest)
}

func (mfa *mfaservice) ValidateOTP(token string, clientID string, clientSecret string) (*mfaClient.ValidateTOTPTokenResponse, error) {
	context, cancel := helpers.InitContext()
	defer cancel()

	validateRequest := &mfaClient.ValidateTOTPTokenRequest{
		Token:        token,
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}

	return mfa.service.ValidateTOTPToken(context, validateRequest)
}

func (mfa *mfaservice) GenerateSecretKey(clientID string) (*mfaClient.TOTPSecretResponse, error) {
	context, cancel := helpers.InitContext()
	defer cancel()

	secretRequest := &mfaClient.GenerateTOTPTokenRequest{
		ClientId: clientID,
	}

	return mfa.service.GenerateSecretKey(context, secretRequest)
}
