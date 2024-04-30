package handlers

import (
	"testing"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces/mocks"
)

func TestPlayerStartConfirmation(t *testing.T) {
	type args struct {
		body   commands.PlayerStartConfirmation
		buffer []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Base case",
			args: args{
				body: commands.PlayerStartConfirmation{
					CharacterID:       0x00001FEF,
					CharacterBodyType: 0x0190,
					Coordinates: commands.Coordinates{
						X: 0x041D,
						Y: 0x0598,
						Z: 0xFFAB,
					},
					DirectionFacing: commands.DirectionFacing{
						Direction: 0x07,
					},
				},
				buffer: []byte{
					0x1B, 0x00, 0x00, 0x1F, 0xEF, 0x00, 0x00, 0x00,
					0x00, 0x01, 0x90, 0x04, 0x1D, 0x05, 0x98, 0xFF,
					0xAB, 0x07, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0x00,
					0x00, 0x00, 0x00, 0x18, 0x00, 0x10, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.CreateClientConnectionWriteBufferMock(t)
			PlayerStartConfirmation(client, tt.args.body)
			client.AssertSentBuffer(tt.args.buffer)
		})
	}
}
