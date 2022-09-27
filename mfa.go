package indicoserviceauth

import (
	mfaClient "github.com/INDICO-INNOVATION/indico_service_auth/client/mfa"
	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/helpers"
	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/iam"
)

func (client *Client) GenerateOTP() (*mfaClient.GenerateOTPTokenResponse, error) {
	context, cancel := helpers.InitContext()
	defer cancel()

	otpRequest := &mfaClient.GenerateOTPTokenRequest{
		ClientId:     iam.Credentials.ClientID,
		ClientSecret: iam.Credentials.ClientSecret,
	}

	return client.MfaService.GenerateOTPToken(context, otpRequest)
}

func (client *Client) ValidateOTP(token string) (*mfaClient.ValidateOTPTokenResponse, error) {
	context, cancel := helpers.InitContext()
	defer cancel()

	validateRequest := &mfaClient.ValidateOTPTokenRequest{
		Token:        token,
		ClientId:     iam.Credentials.ClientID,
		ClientSecret: iam.Credentials.ClientSecret,
	}

	return client.MfaService.ValidateOTPToken(context, validateRequest)
}

func (client *Client) GenerateSecretKey() (*mfaClient.OTPSecretResponse, error) {
	context, cancel := helpers.InitContext()
	defer cancel()

	secretRequest := &mfaClient.GenerateOTPTokenRequest{
		ClientId: iam.Credentials.ClientSecret,
	}

	return client.MfaService.GenerateSecretKey(context, secretRequest)
}
