package login

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	backendLogin "yamul-gateway/backend/services/login"
	"yamul-gateway/internal/transport/multima/commands"
)

const (
	OK int = iota
	INVALID_USER
	INVALID_CREDENTIALS
)

func CheckUserCredentials(username string, password string) (bool, commands.LoginDeniedReason) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	dial, err := grpc.Dial("localhost:8087", opts...)
	if err != nil {
		return false, commands.LoginDeniedReason_CommunicationProblem
	}
	defer dial.Close()
	client := backendLogin.NewLoginServiceClient(dial)
	response, err := client.ValidateLogin(context.Background(), &backendLogin.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return false, commands.LoginDeniedReason_CommunicationProblem
	}
	if response.Value == backendLogin.LoginResponse_valid {
		return true, 0
	}
	return false, commands.LoginDeniedReason_IncorrectUsernamePassword
}
