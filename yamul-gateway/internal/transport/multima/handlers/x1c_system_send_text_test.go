package handlers

import (
	"testing"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces/mocks"
)

func TestSystemSendText(t *testing.T) {
	type args struct {
		body   commands.SystemSendText
		result []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "base case",
			args: args{
				body: commands.SystemSendText{
					Serial: 1,
					Model:  2,
					Type:   3,
					Hue:    4,
					Font:   5,
					Name:   "John Doe",
					Body:   "Hello World!",
				},
				result: []byte{
					0x1c, 0x00, 0x39, 0x00, 0x00, 0x00, 0x01, 0x00,
					0x02, 0x03, 0x00, 0x04, 0x00, 0x05, 0x4a, 0x6f,
					0x68, 0x6e, 0x20, 0x44, 0x6f, 0x65, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x48, 0x65, 0x6c, 0x6c,
					0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x21,
					0x00,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.CreateClientConnectionWriteBufferMock(t)
			SystemSendText(client, tt.args.body)
			client.AssertSentBuffer(tt.args.result)
			client.AssertDeclaredLength(1)
		})
	}
}
