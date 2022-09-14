package mfa

var mfaservice = NewMFAServiceClient(GrpcConnect("localhost:50051"))

func GenerateOTP(clientID string, clientSecret string) (*GenerateTOTPTokenResponse, error) {
	context, cancel := InitContext()
	defer cancel()

	otpRequest := &GenerateTOTPTokenRequest{
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}

	return mfaservice.GenerateTOTPToken(context, otpRequest)
}

func ValidateOTP(token int32, clientID string, clientSecret string) (*ValidateTOTPTokenResponse, error) {
	context, cancel := InitContext()
	defer cancel()

	validateRequest := &ValidateTOTPTokenRequest{
		Token:        token,
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}

	return mfaservice.ValidateTOTPToken(context, validateRequest)
}

func GenerateSecretKey(clientID string) (*TOTPSecretResponse, error) {
	context, cancel := InitContext()
	defer cancel()

	secretRequest := &GenerateTOTPTokenRequest{
		ClientId: clientID,
	}

	return mfaservice.GenerateSecretKey(context, secretRequest)
}
