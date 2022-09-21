package mfa

import (
	"fmt"

	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/helpers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect() *grpc.ClientConn {
	credentials := helpers.ParseEnvironment()

	conn, err := grpc.Dial(credentials.AuthURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("error")
		return nil
	}

	fmt.Printf("grpc is now connected to %s\n", credentials.AuthURI)

	return conn
}
