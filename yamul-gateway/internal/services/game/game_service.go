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

func CreateGameService(connection interfaces.ClientConnection) (interfaces.GameService, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	dial, err := grpc.Dial("localhost:8089", opts...)
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
	result := &gameService{
		dial:              dial,
		client:            client,
		stream:            stream,
		clientConnection:  connection,
		streamLoopEnabled: true,
	}

	go result.streamLoop()

	return result, nil
}

type gameService struct {
	dial              *grpc.ClientConn
	client            backendServices.GameServiceClient
	stream            backendServices.GameService_OpenGameStreamClient
	clientConnection  interfaces.ClientConnection
	streamLoopEnabled bool
}

func (s gameService) Send(_type backendServices.MsgType, message *backendServices.Message) {
	err := s.stream.Send(&backendServices.StreamPackage{
		Type: _type,
		Body: message,
	})
	if err != nil {
		s.clientConnection.KillConnection(err)
		return
	}
}

func (s gameService) Close() {
	s.streamLoopEnabled = false
}

func (s gameService) streamLoop() {
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

func (s gameService) cleanResources() {
	_ = s.dial.Close()
}
