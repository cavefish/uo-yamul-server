package gameEvents

import (
	"reflect"
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
)

func Test_mapToGeneralLightLevel(t *testing.T) {
	type args struct {
		level *services.MsgGeneralLightLevel
	}
	tests := []struct {
		name string
		args args
		want commands.GeneralLightLevel
	}{
		{
			name: "base case",
			args: args{
				level: &services.MsgGeneralLightLevel{
					Level: 42,
				},
			},
			want: commands.GeneralLightLevel{
				Level: 42,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapToCommandGeneralLightLevel(tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapToCommandGeneralLightLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
