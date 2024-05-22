package gameEvents

import (
	"reflect"
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
)

func Test_mapToExtendedStats(t *testing.T) {
	type args struct {
		input *services.MsgExtendedStats
	}
	tests := []struct {
		name string
		args args
		want commands.ExtendedStats
	}{
		{
			name: "dead case",
			args: args{
				input: &services.MsgExtendedStats{
					Msg: &services.MsgExtendedStats_Dead{Dead: &services.MsgExtendedStats_MsgExtendedStats_Dead{
						Id:     &services.ObjectId{Value: 42},
						IsDead: true,
					}},
				},
			},
			want: commands.ExtendedStats{
				Type:     commands.ExtendedStats_Dead,
				ObjectID: 42,
				StrLock:  false,
				DexLock:  false,
				IntLock:  false,
				IsDead:   true,
			},
		},
		{
			name: "false dead case",
			args: args{
				input: &services.MsgExtendedStats{
					Msg: &services.MsgExtendedStats_Dead{Dead: &services.MsgExtendedStats_MsgExtendedStats_Dead{
						Id:     &services.ObjectId{Value: 421},
						IsDead: false,
					}},
				},
			},
			want: commands.ExtendedStats{
				Type:     commands.ExtendedStats_Dead,
				ObjectID: 421,
				StrLock:  false,
				DexLock:  false,
				IntLock:  false,
				IsDead:   false,
			},
		},
		{
			name: "lock case",
			args: args{
				input: &services.MsgExtendedStats{
					Msg: &services.MsgExtendedStats_Lock{Lock: &services.MsgExtendedStats_MsgExtendedStats_AttributeLock{
						Id:      &services.ObjectId{Value: 42},
						StrLock: true,
						DexLock: true,
						IntLock: true,
					}},
				},
			},
			want: commands.ExtendedStats{
				Type:     commands.ExtendedStats_AttributeLock,
				ObjectID: 42,
				StrLock:  true,
				DexLock:  true,
				IntLock:  true,
				IsDead:   false,
			},
		},
		{
			name: "lock none case",
			args: args{
				input: &services.MsgExtendedStats{
					Msg: &services.MsgExtendedStats_Lock{Lock: &services.MsgExtendedStats_MsgExtendedStats_AttributeLock{
						Id:      &services.ObjectId{Value: 4200},
						StrLock: false,
						DexLock: false,
						IntLock: false,
					}},
				},
			},
			want: commands.ExtendedStats{
				Type:     commands.ExtendedStats_AttributeLock,
				ObjectID: 4200,
				StrLock:  false,
				DexLock:  false,
				IntLock:  false,
				IsDead:   false,
			},
		},
		{
			name: "lock str case",
			args: args{
				input: &services.MsgExtendedStats{
					Msg: &services.MsgExtendedStats_Lock{Lock: &services.MsgExtendedStats_MsgExtendedStats_AttributeLock{
						Id:      &services.ObjectId{Value: 4200},
						StrLock: true,
						DexLock: false,
						IntLock: false,
					}},
				},
			},
			want: commands.ExtendedStats{
				Type:     commands.ExtendedStats_AttributeLock,
				ObjectID: 4200,
				StrLock:  true,
				DexLock:  false,
				IntLock:  false,
				IsDead:   false,
			},
		},
		{
			name: "lock dex case",
			args: args{
				input: &services.MsgExtendedStats{
					Msg: &services.MsgExtendedStats_Lock{Lock: &services.MsgExtendedStats_MsgExtendedStats_AttributeLock{
						Id:      &services.ObjectId{Value: 4200},
						StrLock: false,
						DexLock: true,
						IntLock: false,
					}},
				},
			},
			want: commands.ExtendedStats{
				Type:     commands.ExtendedStats_AttributeLock,
				ObjectID: 4200,
				StrLock:  false,
				DexLock:  true,
				IntLock:  false,
				IsDead:   false,
			},
		},
		{
			name: "lock int case",
			args: args{
				input: &services.MsgExtendedStats{
					Msg: &services.MsgExtendedStats_Lock{Lock: &services.MsgExtendedStats_MsgExtendedStats_AttributeLock{
						Id:      &services.ObjectId{Value: 4200},
						StrLock: false,
						DexLock: false,
						IntLock: true,
					}},
				},
			},
			want: commands.ExtendedStats{
				Type:     commands.ExtendedStats_AttributeLock,
				ObjectID: 4200,
				StrLock:  false,
				DexLock:  false,
				IntLock:  true,
				IsDead:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapToCommandExtendedStats(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapToCommandExtendedStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
