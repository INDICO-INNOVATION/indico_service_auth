package iam

import (
	"log"

	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/helpers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Credentials, authServer = helpers.ParseEnvironment()

func Connect() *grpc.ClientConn {
	conn, err := grpc.Dial(authServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error to conenct to auth server due to %s", err.Error())
	}

	return conn
}
