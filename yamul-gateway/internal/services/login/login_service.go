package login

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	backendLogin "yamul-gateway/backend/services/login"
	"yamul-gateway/internal/transport/multima/commands"
)

type LoginService struct {
	dial   *grpc.ClientConn
	client backendLogin.LoginServiceClient
}

var Service *LoginService

func Setup() error {
	service, err := newLoginService()
	Service = service
	return err
}

func Close() {
	Service.close()
}

func newLoginService() (*LoginService, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	dial, err := grpc.Dial("localhost:8087", opts...)
	if err != nil {
		return nil, err
	}
	return &LoginService{
		dial:   dial,
		client: backendLogin.NewLoginServiceClient(dial),
	}, nil
}

func (s LoginService) close() {
	err := s.dial.Close()
	if err != nil {
		return
	}
}

func (s LoginService) CheckUserCredentials(username string, password string) (bool, commands.LoginDeniedReason) {
	response, err := s.client.ValidateLogin(context.Background(), &backendLogin.LoginRequest{
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
