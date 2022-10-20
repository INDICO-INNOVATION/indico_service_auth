package iam

import (
	"crypto/tls"
	"crypto/x509"
	"log"

	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/helpers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var Credentials, authServer = helpers.ParseEnvironment()

func Connect() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithAuthority(authServer))

	insecure := false
	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			log.Fatalf(err.Error())
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	conn, err := grpc.Dial(authServer, opts...)
	if err != nil {
		log.Fatalf("error to conenct to auth server due to %s", err.Error())
	}

	return conn
}
