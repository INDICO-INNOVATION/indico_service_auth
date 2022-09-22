package mfa

import (
	"fmt"

	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/helpers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect() *grpc.ClientConn {
	// TODO: use credentials
	_, authServer := helpers.ParseEnvironment()

	conn, err := grpc.Dial(authServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("error")
		return nil
	}

	return conn
}
