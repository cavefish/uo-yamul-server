package handlers

import (
	"testing"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces/mocks"
)

func TestGeneralLightLevel(t *testing.T) {
	type args struct {
		output  []byte
		command commands.GeneralLightLevel
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Base case",
			args: args{
				output:  []byte{0x4f, 0x18},
				command: commands.GeneralLightLevel{Level: 0x18},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.CreateClientConnectionWriteBufferMock(t)
			GeneralLightLevel(client, tt.args.command)
			client.AssertSentBuffer(tt.args.output)
		})
	}
}
