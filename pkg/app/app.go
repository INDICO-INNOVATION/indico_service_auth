package app

import (
	grpcHelper "github.com/INDICO-INNOVATION/indico_service_auth/pkg/grpc"
	"google.golang.org/grpc"
)

var Inst *Application

type Application struct {
	GrpcServer *grpc.ClientConn
}

func ApplicationInit(mfaServer string) {
	Inst = &Application{
		GrpcServer: grpcHelper.Connect(mfaServer),
	}
}
