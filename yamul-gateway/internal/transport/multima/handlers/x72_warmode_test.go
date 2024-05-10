package handlers

import (
	"testing"
	"yamul-gateway/internal/interfaces/mocks"
)

func TestWarmode(t *testing.T) {
	type args struct {
		output    []byte
		isWarmode bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "negative case",
			args: args{
				output: []byte{
					0x72, 0x00, 0x00, 0x32, 0x00,
				},
				isWarmode: false,
			},
		},
		{
			name: "positive case",
			args: args{
				output: []byte{
					0x72, 0x01, 0x00, 0x32, 0x00,
				},
				isWarmode: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.CreateClientConnectionWriteBufferMock(t)
			Warmode(client, tt.args.isWarmode)
			client.AssertSentBuffer(tt.args.output)
		})
	}
}
