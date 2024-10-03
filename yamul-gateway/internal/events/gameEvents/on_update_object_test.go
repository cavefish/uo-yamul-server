package gameEvents

import (
	"math"
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

func Test_toCommandUpdateObject1(t *testing.T) {
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
					Id: &services.ObjectId{
						Value: 1,
					},
					GraphicId:      2,
					XLoc:           3,
					YLoc:           4,
					ZLoc:           5,
					Direction:      6,
					Hue:            7,
					Flags:          []services.Flags{8},
					NotorietyFlags: []services.Notoriety{9},
				},
			},
			want: commands.UpdateObject{
				Serial:        1,
				GraphicId:     2,
				XLoc:          3,
				YLoc:          4,
				ZLoc:          5,
				Direction:     6,
				Hue:           7,
				Flags:         8,
				NotorietyFlag: 9,
			},
		},
		{
			name: "max values case",
			args: args{
				object: &services.MsgUpdateObject{
					Id: &services.ObjectId{
						Value: math.MaxUint32,
					},
					GraphicId:      math.MaxUint16,
					XLoc:           math.MaxUint16,
					YLoc:           math.MaxUint16,
					ZLoc:           0xFF,
					Direction:      0xFF,
					Hue:            math.MaxUint16,
					Flags:          []services.Flags{0xFF},
					NotorietyFlags: []services.Notoriety{0xFF},
				},
			},
			want: commands.UpdateObject{
				Serial:        math.MaxUint32,
				GraphicId:     math.MaxUint16,
				XLoc:          math.MaxUint16,
				YLoc:          math.MaxUint16,
				ZLoc:          0xFF,
				Direction:     0xFF,
				Hue:           math.MaxUint16,
				Flags:         0xFF,
				NotorietyFlag: 0xFF,
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
