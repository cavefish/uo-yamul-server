package handlers

import (
	"testing"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces/mocks"
)

func TestExtendedStats(t *testing.T) {
	type args struct {
		output  []byte
		command commands.ExtendedStats
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Dead case",
			args: args{
				output: []byte{
					0xBF, 0x00, 0x0B, 0x00, 0x19, 0x00, 0x00, 0x00,
					0x1F, 0xEF, 0xFF,
				},
				command: commands.ExtendedStats{
					Type:     commands.ExtendedStats_Dead,
					ObjectID: 0x00001fef,
					IsDead:   true,
				},
			},
		},
		{
			name: "AttributeLock case",
			args: args{
				output: []byte{
					0xBF, 0x00, 0x0C, 0x00, 0x19, 0x02, 0x00, 0x00,
					0x1F, 0xEF, 0x00, 0b00111111,
				},
				command: commands.ExtendedStats{
					Type:     commands.ExtendedStats_AttributeLock,
					ObjectID: 0x00001fef,
					StrLock:  true,
					DexLock:  true,
					IntLock:  true,
				},
			},
		},
		{
			name: "real case",
			args: args{
				output: []byte{
					0xBF, 0x00, 0x0C, 0x00, 0x19, 0x02, 0x00, 0x00,
					0x1F, 0xEF, 0x00, 0x00,
				},
				command: commands.ExtendedStats{
					Type:     commands.ExtendedStats_AttributeLock,
					ObjectID: 0x00001fef,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.CreateClientConnectionWriteBufferMock(t)
			ExtendedStats(client, tt.args.command)
			client.AssertSentBuffer(tt.args.output)
		})
	}
}
