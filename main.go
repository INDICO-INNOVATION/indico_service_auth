package go_mfaservice

import (
	mfaClient "github.com/INDICO-INNOVATION/indico_service_auth/client/mfa"
	"github.com/INDICO-INNOVATION/indico_service_auth/config"
	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/app"
)

type MfaService interface {
	GenerateOTP(clientID string, clientSecret string) (*mfaClient.GenerateTOTPTokenResponse, error)
	ValidateOTP(token int32, clientID string, clientSecret string) (*mfaClient.ValidateTOTPTokenResponse, error)
	GenerateSecretKey(clientID string) (*mfaClient.TOTPSecretResponse, error)
}

type mfaservice struct {
	service mfaClient.MFAServiceClient
}

func main() {
	app.ApplicationInit(config.New())
}

func NewMFAService() MfaService {
	return &mfaservice{
		service: mfaClient.NewMFAServiceClient(app.Inst.GrpcServer),
	}
}
