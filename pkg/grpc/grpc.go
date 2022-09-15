package mfa

import (
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO: Change uri when it will be available in prod
const AUTH_SERVER = "localhost:50051"

func Connect() *grpc.ClientConn {
	var addr string
	if os.Getenv("AUTH_SERVER") != "" {
		addr = os.Getenv("AUTH_SERVER")
	} else {
		addr = AUTH_SERVER
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("error")
		return nil
	}

	fmt.Printf("grpc is now connected to %s", addr)

	return conn
}
