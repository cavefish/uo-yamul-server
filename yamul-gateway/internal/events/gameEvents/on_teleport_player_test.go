package gameEvents

import (
	"math"
	"reflect"
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
)

func Test_toCommandTeleportPlayer(t *testing.T) {
	type args struct {
		player *services.MsgTeleportPlayer
	}
	tests := []struct {
		name string
		args args
		want commands.TeleportPlayer
	}{
		{
			name: "base case",
			args: args{player: &services.MsgTeleportPlayer{
				Id:     &services.ObjectId{Value: 0},
				Status: []services.MsgTeleportPlayer_PlayerStatus{},
				Coordinates: &services.Coordinate{
					XLoc: 0,
					YLoc: 0,
					ZLoc: 0,
				},
				Direction: 0,
			}},
			want: commands.TeleportPlayer{
				Serial:    0,
				Status:    0,
				XLoc:      0,
				YLoc:      0,
				Direction: 0,
				ZLoc:      0,
			},
		},
		{
			name: "with some values",
			args: args{player: &services.MsgTeleportPlayer{
				Id: &services.ObjectId{Value: 1},
				Status: []services.MsgTeleportPlayer_PlayerStatus{
					services.MsgTeleportPlayer_PlayerStatus_normal,
				},
				Coordinates: &services.Coordinate{
					XLoc: 2,
					YLoc: 3,
					ZLoc: 6,
				},
				Direction: services.ObjectDirection_left,
			}},
			want: commands.TeleportPlayer{
				Serial:    1,
				Status:    0,
				XLoc:      2,
				YLoc:      3,
				Direction: 5,
				ZLoc:      6,
			},
		},
		{
			name: "with upper limit values values",
			args: args{player: &services.MsgTeleportPlayer{
				Id: &services.ObjectId{Value: 100},
				Status: []services.MsgTeleportPlayer_PlayerStatus{
					services.MsgTeleportPlayer_PlayerStatus_normal,
					services.MsgTeleportPlayer_PlayerStatus_poisoned,
					services.MsgTeleportPlayer_PlayerStatus_warMode,
					services.MsgTeleportPlayer_PlayerStatus_canAlterPaperDoll,
				},
				Coordinates: &services.Coordinate{
					XLoc: math.MaxUint16,
					YLoc: math.MaxUint16,
					ZLoc: math.MaxInt8,
				},
				Direction: services.ObjectDirection_running | services.ObjectDirection_down,
			}},
			want: commands.TeleportPlayer{
				Serial:    100,
				Status:    0x46,
				XLoc:      math.MaxUint16,
				YLoc:      math.MaxUint16,
				Direction: 128 | 3,
				ZLoc:      math.MaxInt8,
			},
		},
		{
			name: "with lower limit values",
			args: args{player: &services.MsgTeleportPlayer{
				Id:     &services.ObjectId{Value: 0},
				Status: []services.MsgTeleportPlayer_PlayerStatus{},
				Coordinates: &services.Coordinate{
					XLoc: 0,
					YLoc: 0,
					ZLoc: math.MinInt8,
				},
				Direction: services.ObjectDirection_north,
			}},
			want: commands.TeleportPlayer{
				Serial:    0,
				Status:    0,
				XLoc:      0,
				YLoc:      0,
				Direction: 0,
				ZLoc:      math.MinInt8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toCommandTeleportPlayer(tt.args.player); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toCommandTeleportPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
