package indicoserviceauth

import (
	"context"
	"log"

	authClient "github.com/INDICO-INNOVATION/indico_service_auth/client/auth"
	mfaClient "github.com/INDICO-INNOVATION/indico_service_auth/client/mfa"
	"github.com/INDICO-INNOVATION/indico_service_auth/pkg/iam"
)

type Client struct {
	authService authClient.AuthServiceClient
	mfaService  mfaClient.MFAServiceClient
}

func generateJWT(context context.Context, authservice authClient.AuthServiceClient, scope string) *authClient.GenerateTokenResponse {
	request := &authClient.GenerateTokenRequest{
		ClientId:     iam.Credentials.ClientID,
		ClientSecret: iam.Credentials.ClientSecret,
		Scope:        scope,
		PrivateKey:   iam.Credentials.PrivateKey,
		Type:         iam.Credentials.Type,
	}

	jwt, err := authservice.GenerateToken(context, request)
	if err != nil {
		log.Fatalf("error to generate jwt token: %s", err.Error())
	}
	return jwt
}

func login(context context.Context, authservice authClient.AuthServiceClient, accessToken string, scope string) *authClient.AuthResponse {
	request := &authClient.AuthRequest{
		AccessToken: accessToken,
		PrivateKey:  iam.Credentials.PrivateKey,
	}

	authenticated, err := authservice.Authenticate(context, request)
	if err != nil {
		log.Fatalf("error to authenticate: %s", err.Error())
	}

	return authenticated
}

func NewClient(context context.Context, scope string) (*Client, error) {
	conn := iam.Connect()

	client := &Client{
		authService: authClient.NewAuthServiceClient(conn),
		mfaService:  mfaClient.NewMFAServiceClient(conn),
	}

	token := generateJWT(context, client.authService, scope)
	authenticated := login(context, client.authService, token.AccessToken, scope)
	if !authenticated.Authenticated {
		log.Fatalln("invalid credentials")
	}

	return client, nil
}
