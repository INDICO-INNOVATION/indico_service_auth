package go_mfaservice

import (
	mfaClient "github.com/INDICO-INNOVATION/indico_service_auth/client/mfa"
	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/app"
)

func NewMFAService(mfaServer string) MfaService {
	app.ApplicationInit(mfaServer)
	return &mfaservice{
		service: mfaClient.NewMFAServiceClient(app.Inst.GrpcServer),
	}
}
