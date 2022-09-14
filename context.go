package mfa

import (
	"context"
	"time"
)

func InitContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(60*time.Second))
}
