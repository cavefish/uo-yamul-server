package game

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	backendServices "yamul-gateway/backend/services"
	"yamul-gateway/internal/interfaces"
	servicesCommon "yamul-gateway/internal/services/common"
	"yamul-gateway/internal/services/game/messages"
)

type GameService struct {
	dial              *grpc.ClientConn
	client            backendServices.GameServiceClient
	stream            backendServices.GameService_OpenGameStreamClient
	clientConnection  interfaces.ClientConnection
	streamLoopEnabled bool
}

func (s GameService) Close() {
	s.streamLoopEnabled = false
}

func (s GameService) streamLoop() {
	defer s.cleanResources()
	for s.streamLoopEnabled {
		msg, err := s.stream.Recv()
		if err != nil {
			s.clientConnection.KillConnection(err)
			return
		}
		processor, ok := messages.Processors[msg.Type]
		if !ok {
			s.clientConnection.KillConnection(fmt.Errorf("Unknown message type %d: %x", msg.Type, msg.Body))
		}
		processor.Accept(s.clientConnection, msg)
	}
}

func (s GameService) cleanResources() {
	_ = s.dial.Close()
}

func NewGameService(connection interfaces.ClientConnection) (*GameService, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	dial, err := grpc.Dial("localhost:8088", opts...)
	if err != nil {
		return nil, err
	}

	client := backendServices.NewGameServiceClient(dial)
	ctx := servicesCommon.GetAuthenticatedContext(context.Background(), connection.GetLoginDetails())
	stream, err := client.OpenGameStream(ctx)
	if err != nil {
		_ = dial.Close()
		return nil, err
	}
	result := &GameService{
		dial:              dial,
		client:            client,
		stream:            stream,
		clientConnection:  connection,
		streamLoopEnabled: true,
	}

	go result.streamLoop()

	return result, nil
}
