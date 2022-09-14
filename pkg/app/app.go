package app

import (
	"github.com/INDICO-INNOVATION/indico_service_auth/config"
	grpcHelper "github.com/INDICO-INNOVATION/indico_service_auth/pkg/grpc"
	"google.golang.org/grpc"
)

var Inst *Application

type Application struct {
	Config     *config.Config
	GrpcServer *grpc.ClientConn
}

func ApplicationInit(configs *config.Config) {
	Inst = &Application{
		Config:     configs,
		GrpcServer: grpcHelper.Connect(configs.MFAServer),
	}
}
