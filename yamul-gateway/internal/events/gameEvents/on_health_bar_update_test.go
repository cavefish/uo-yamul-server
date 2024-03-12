package gameEvents

import (
	"math"
	"reflect"
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
)

func Test_mapToHealthBarUpdate(t *testing.T) {
	type args struct {
		body *services.MsgHealthBar
	}
	tests := []struct {
		name string
		args args
		want commands.HealthBarUpdate
	}{
		{
			name: "base case",
			args: args{body: &services.MsgHealthBar{
				Id:     &services.ObjectId{Value: 0},
				Values: []*services.MsgHealthBar_Values{},
			}},
			want: commands.HealthBarUpdate{
				Serial: 0,
				Values: []commands.HealthBarUpdateValues{},
			},
		},
		{
			name: "edge case",
			args: args{body: &services.MsgHealthBar{
				Id: &services.ObjectId{Value: math.MaxUint32},
				Values: []*services.MsgHealthBar_Values{
					{
						Type:    services.MsgHealthBar_Values_GREEN,
						Enabled: false,
					},
					{
						Type:    services.MsgHealthBar_Values_YELLOW,
						Enabled: true,
					},
				},
			}},
			want: commands.HealthBarUpdate{
				Serial: 0,
				Values: []commands.HealthBarUpdateValues{
					{
						Type:    commands.HealthBarUpdateValues_Green,
						Enabled: false,
					},
					{
						Type:    commands.HealthBarUpdateValues_Yellow,
						Enabled: true,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapToHealthBarUpdate(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapToHealthBarUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
