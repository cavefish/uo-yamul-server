package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces/mocks"
)

func Test_moveAckReadBuffer(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want commands.MoveAck
	}{
		{
			name: "base case",
			args: args{
				input: []byte{0xab, 0xcd},
			},
			want: commands.MoveAck{
				Sequence: 0xab,
				Status:   0xcd,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.CreateClientConnectionReadBufferMock(t, tt.args.input)
			assert.Equalf(t, tt.want, moveAckReadBuffer(client), "moveAckReadBuffer(%v)", client)
			client.AssertBufferConsumed()
		})
	}
}
