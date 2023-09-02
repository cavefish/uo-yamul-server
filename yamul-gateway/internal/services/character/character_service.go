package character

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	backendServices "yamul-gateway/backend/services"
	servicesCommon "yamul-gateway/internal/services/common"
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/internal/transport/multima/connection"
)

type CharacterService struct {
	dial       *grpc.ClientConn
	client     backendServices.CharacterServiceClient
	connection *connection.ClientConnection
}

func (s CharacterService) Close() {
	_ = s.dial.Close()
}

func (s CharacterService) GetCharacters() ([]commands.CharacterLogin, int, error) {
	ctx := servicesCommon.GetAuthenticatedContext(context.Background(), s.connection.LoginDetails)
	response, err := s.client.GetCharacterList(ctx, &backendServices.Empty{})
	if err != nil {
		return nil, 0, err
	}
	lastValidCharacter := len(response.CharacterLogins) - 1
	result := make([]commands.CharacterLogin, 5)

	for i := 0; i <= lastValidCharacter; i++ {
		result[i].Name = response.CharacterLogins[i].Username
		result[i].Password = response.CharacterLogins[i].Password
	}
	for i := lastValidCharacter + 1; i < len(result); i++ {
		result[i].Name = "unused"
		result[i].Password = ""
	}

	return result, lastValidCharacter, nil
}

func NewCharacterService(connection *connection.ClientConnection) (*CharacterService, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	dial, err := grpc.Dial("localhost:8088", opts...)
	if err != nil {
		return nil, err
	}
	return &CharacterService{
		dial:       dial,
		client:     backendServices.NewCharacterServiceClient(dial),
		connection: connection,
	}, nil
}
