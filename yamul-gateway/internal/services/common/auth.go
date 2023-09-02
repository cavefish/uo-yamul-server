package common

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"yamul-gateway/internal/transport/multima/connection"
)

func GetAuthenticatedContext(ctx context.Context, loginDetails connection.LoginDetails) context.Context {
	token := fmt.Sprintf("%s %s", loginDetails.Username, loginDetails.Password)
	return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"x-auth-key": token}))
}
