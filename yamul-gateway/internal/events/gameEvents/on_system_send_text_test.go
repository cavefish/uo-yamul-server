package gameEvents

import (
	"reflect"
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
)

func Test_mapToCommandSystemSendText(t *testing.T) {
	type args struct {
		text *services.MsgSystemSendText
	}
	tests := []struct {
		name string
		args args
		want commands.SystemSendText
	}{
		{
			name: "base case",
			args: args{
				text: &services.MsgSystemSendText{
					Id: &services.ObjectId{
						Value: 1,
					},
					Model: 2,
					Type:  3,
					Hue:   4,
					Font:  5,
					Name:  "Jane Doe",
					Body:  "FooBar",
				},
			},
			want: commands.SystemSendText{
				Serial: 1,
				Model:  2,
				Type:   3,
				Hue:    4,
				Font:   5,
				Name:   "Jane Doe",
				Body:   "FooBar",
			},
		},
		{
			name: "System message case",
			args: args{
				text: &services.MsgSystemSendText{
					Id:    nil,
					Model: 0,
					Type:  1,
					Hue:   2,
					Font:  3,
					Name:  "John Doe",
					Body:  "Hello World!",
				},
			},
			want: commands.SystemSendText{
				Serial: 0xFFFFFFFF,
				Model:  0xFFFF,
				Type:   1,
				Hue:    2,
				Font:   3,
				Name:   "John Doe",
				Body:   "Hello World!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapToCommandSystemSendText(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapToCommandSystemSendText() = %v, want %v", got, tt.want)
			}
		})
	}
}
