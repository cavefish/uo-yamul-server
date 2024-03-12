package gameEvents

import (
	"reflect"
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
)

func Test_toCommandMapChange(t *testing.T) {
	type args struct {
		body *services.MsgMapChange
	}
	tests := []struct {
		name string
		args args
		want commands.MapChange
	}{
		{
			name: "min value",
			args: args{body: &services.MsgMapChange{MapId: 0}},
			want: commands.MapChange{MapId: 0},
		},
		{
			name: "max value",
			args: args{body: &services.MsgMapChange{MapId: 0xff}},
			want: commands.MapChange{MapId: 0xff},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toCommandMapChange(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toCommandMapChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
