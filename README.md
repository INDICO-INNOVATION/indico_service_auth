# Indico Innovation MFA Client

[![pkg-img]][pkg-url]

## Rationale

We were looking for a simple way to use Indico Innovation MFA service in Indico Innovation IAM GRPC server. We will have to develop a new GRPC client every time a client application need to connect to IAM server. 

So we have built this MFA Client to be imported by any Golang application instead to deploy it as a microservice avoiding to make http requests to work.

## Dependencies

* MFA protobuffers

## Features

* Connect to MFA Service in IAM GRPC server.

## Install
Go version 1.14+
```
go get github.com/INDICO-INNOVATION/indico_service_auth
```


```bash
# Example
import (
    mfaservice "github.com/INDICO-INNOVATION/indico_service_auth"
)

mfaService := mfaservice.NewMFAService("<iam_grpc_server_host>")
response, err := mfaService.GenerateSecretKey("<client_id>")

```

## Documentation

See in [docs](https://pkg.go.dev/github.com/INDICO-INNOVATION/indico_auth_service).

## License

[GNU General Public License v3.0](./LICENSE).
