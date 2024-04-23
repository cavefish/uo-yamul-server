package gameEvents

import (
	"reflect"
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
)

func Test_toCommandUpdateObject(t *testing.T) {
	type args struct {
		object *services.MsgUpdateObject
	}
	tests := []struct {
		name string
		args args
		want commands.UpdateObject
	}{
		{
			name: "base case",
			args: args{
				object: &services.MsgUpdateObject{
					Id: &services.ObjectId{Value: 42},
				},
			},
			want: commands.UpdateObject{
				Serial:        42,
				GraphicId:     0,
				XLoc:          0,
				YLoc:          0,
				ZLoc:          0,
				Direction:     0,
				Hue:           0,
				Flags:         0,
				NotorietyFlag: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toCommandUpdateObject(tt.args.object); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toCommandUpdateObject() = %v, want %v", got, tt.want)
			}
		})
	}
}
