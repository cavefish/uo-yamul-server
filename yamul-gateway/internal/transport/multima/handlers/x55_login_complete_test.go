package handlers

import (
	"testing"
	"yamul-gateway/internal/interfaces/mocks"
)

func TestLoginComplete(t *testing.T) {
	type args struct {
		output  []byte
		nothing any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "only case",
			args: args{
				output:  []byte{0x55},
				nothing: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.CreateClientConnectionWriteBufferMock(t)
			LoginComplete(client, tt.args.nothing)
			client.AssertSentBuffer(tt.args.output)
		})
	}
}
