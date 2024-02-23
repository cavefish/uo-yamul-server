package common

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"yamul-gateway/internal/dtos"
)

func GetAuthenticatedContext(ctx context.Context, loginDetails *dtos.LoginDetails) context.Context {
	token := fmt.Sprintf("%s %s", loginDetails.Username, loginDetails.Password)
	return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"x-auth-key": token}))
}
