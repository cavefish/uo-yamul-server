package gameEvents

import (
	"reflect"
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
)

func Test_mapToSkillUpdateServerCommand(t *testing.T) {
	type args struct {
		input *services.MsgSkillUpdateServer
	}
	tests := []struct {
		name string
		args args
		want commands.SkillUpdateServer
	}{
		{
			name: "basic case",
			args: args{
				input: &services.MsgSkillUpdateServer{
					Type: services.MsgSkillUpdateType_basic,
					Skills: []*services.MsgSkillUpdateServer_MsgSkillUpdateSkills{
						{
							SkillId:   1,
							Value:     2,
							BaseValue: 3,
							Status:    4,
							MaxValue:  5,
						},
						{
							SkillId:   6,
							Value:     7,
							BaseValue: 8,
							Status:    9,
							MaxValue:  10,
						},
					},
				},
			},
			want: commands.SkillUpdateServer{
				Type: 0,
				Skills: []*commands.SkillUpdateServerSkill{
					{
						Id:        1,
						Value:     2,
						BaseValue: 3,
						MaxValue:  5,
						Status:    4,
					},
					{
						Id:        6,
						Value:     7,
						BaseValue: 8,
						MaxValue:  10,
						Status:    9,
					},
				},
			},
		},
		{
			name: "Other case",
			args: args{
				input: &services.MsgSkillUpdateServer{
					Type: services.MsgSkillUpdateType_updateSkillCap,
					Skills: []*services.MsgSkillUpdateServer_MsgSkillUpdateSkills{
						{
							SkillId:   1,
							Value:     2,
							BaseValue: 3,
							Status:    4,
							MaxValue:  5,
						},
					},
				},
			},
			want: commands.SkillUpdateServer{
				Type: 0xDF,
				Skills: []*commands.SkillUpdateServerSkill{
					{
						Id:        1,
						Value:     2,
						BaseValue: 3,
						MaxValue:  5,
						Status:    4,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapToSkillUpdateServerCommand(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapToSkillUpdateServerCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
