package mfa

import (
	"google.golang.org/grpc"
)

var Inst *Application

type Application struct {
	Config     *Config
	GrpcServer *grpc.ClientConn
}

func ApplicationInit(configs *Config) {
	Inst = &Application{
		Config:     configs,
		GrpcServer: GrpcConnect(configs.MFAServer),
	}
}
