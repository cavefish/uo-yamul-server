package login

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	backendServices "yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
)

type LoginService struct {
	dial   *grpc.ClientConn
	client backendServices.LoginServiceClient
}

var service *LoginService

func Setup() error {
	var err error
	service, err = newLoginService()
	return err
}

func Close() {
	service.close()
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
		client: backendServices.NewLoginServiceClient(dial),
	}, nil
}

func (s LoginService) close() {
	_ = s.dial.Close()
}

func (s LoginService) CheckUserCredentials(username string, password string) (bool, commands.LoginDeniedReason) {
	response, err := s.client.ValidateLogin(context.Background(), &backendServices.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return false, commands.LoginDeniedReason_CommunicationProblem
	}
	if response.Value == backendServices.LoginResponse_valid {
		return true, 0
	}
	return false, commands.LoginDeniedReason_IncorrectUsernamePassword
}
