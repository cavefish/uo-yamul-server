package gameEvents

import (
	"math"
	"reflect"
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
)

func Test_toCommandPlayMusic(t *testing.T) {
	type args struct {
		music *services.MsgPlayMusic
	}
	tests := []struct {
		name string
		args args
		want commands.PlayMusic
	}{
		{
			name: "min values",
			args: args{music: &services.MsgPlayMusic{MusicId: 0}},
			want: commands.PlayMusic{MusicId: 0},
		},
		{
			name: "max values",
			args: args{music: &services.MsgPlayMusic{MusicId: math.MaxUint16}},
			want: commands.PlayMusic{MusicId: math.MaxUint16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toCommandPlayMusic(tt.args.music); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toCommandPlayMusic() = %v, want %v", got, tt.want)
			}
		})
	}
}
