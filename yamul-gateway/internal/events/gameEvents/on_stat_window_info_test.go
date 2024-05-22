package gameEvents

import (
	"reflect"
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
)

func Test_mapToCommandStatWindowInfo(t *testing.T) {
	type args struct {
		info *services.MsgStatWindowInfo
	}
	tests := []struct {
		name string
		args args
		want commands.StatWindowInfo
	}{
		{
			name: "Full case",
			args: args{
				info: &services.MsgStatWindowInfo{
					CharacterID: &services.ObjectId{
						Value: 42000,
					},
					CharacterName:    "test",
					HitPointsCurrent: 1,
					HitPointsMax:     2,
					FlagNameAllowed:  true,
					FlagDisplay:      3,
					Level2: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2{
						Gender:             4,
						Strength:           5,
						Intelligence:       6,
						StaminaCurrent:     7,
						StaminaMax:         8,
						ManaCurrent:        9,
						ManaMax:            10,
						Gold:               11,
						ResistancePhysical: 12,
						WeightCurrent:      13,
						Level3: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3{
							StatsCap: 14,
							Level4: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3_MsgStatWindowInfoLevel4{
								FollowersCurrent: 15,
								FollowersMax:     16,
								Level5: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3_MsgStatWindowInfoLevel4_MsgStatWindowInfoLevel5{
									ResistanceFire:   17,
									ResistanceCold:   18,
									ResistancePoison: 19,
									ResistanceEnergy: 20,
									Luck:             21,
									DamageMin:        22,
									DamageMax:        23,
									TithingPoints:    24,
									Level6: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3_MsgStatWindowInfoLevel4_MsgStatWindowInfoLevel5_MsgStatWindowInfoLevel6{
										WeightMax: 25,
										Race:      26,
										Level7: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3_MsgStatWindowInfoLevel4_MsgStatWindowInfoLevel5_MsgStatWindowInfoLevel6_MsgStatWindowInfoLevel7{
											ResistancePhysicalMax:        27,
											ResistanceFireMax:            28,
											ResistanceColdMax:            29,
											ResistancePoisonMax:          30,
											ResistanceEnergyMax:          31,
											DefenseChanceIncreaseCurrent: 32,
											DefenseChanceIncreaseMax:     33,
											HitChanceIncrease:            34,
											SwingSpeedIncrease:           35,
											DamageIncrease:               36,
											LowerReagentCost:             37,
											SpellDamageIncrease:          38,
											FasterCastRecovery:           39,
											FasterCasting:                40,
											LowerManaCost:                41,
										},
									},
								},
							},
						},
					},
				},
			},
			want: commands.StatWindowInfo{
				CharacterID:                  42000,
				CharacterName:                "test",
				HitPointsCurrent:             1,
				HitPointsMax:                 2,
				FlagNameAllowed:              true,
				FlagDisplay:                  3,
				Gender:                       4,
				Strength:                     5,
				Intelligence:                 6,
				StaminaCurrent:               7,
				StaminaMax:                   8,
				ManaCurrent:                  9,
				ManaMax:                      10,
				Gold:                         11,
				ResistancePhysical:           12,
				WeightCurrent:                13,
				WeightMax:                    25,
				Race:                         26,
				StatsCap:                     14,
				FollowersCurrent:             15,
				FollowersMax:                 16,
				ResistanceFire:               17,
				ResistanceCold:               18,
				ResistancePoison:             19,
				ResistanceEnergy:             20,
				Luck:                         21,
				DamageMin:                    22,
				DamageMax:                    23,
				TithingPoints:                24,
				ResistancePhysicalMax:        27,
				ResistanceFireMax:            28,
				ResistanceColdMax:            29,
				ResistancePoisonMax:          30,
				ResistanceEnergyMax:          31,
				DefenseChanceIncreaseCurrent: 32,
				DefenseChanceIncreaseMax:     33,
				HitChanceIncrease:            34,
				SwingSpeedIncrease:           35,
				DamageIncrease:               36,
				LowerReagentCost:             37,
				SpellDamageIncrease:          38,
				FasterCastRecovery:           39,
				FasterCasting:                40,
				LowerManaCost:                41,
			},
		},
		{
			name: "level6 case",
			args: args{
				info: &services.MsgStatWindowInfo{
					CharacterID: &services.ObjectId{
						Value: 42000,
					},
					CharacterName:    "test",
					HitPointsCurrent: 1,
					HitPointsMax:     2,
					FlagNameAllowed:  true,
					FlagDisplay:      3,
					Level2: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2{
						Gender:             4,
						Strength:           5,
						Intelligence:       6,
						StaminaCurrent:     7,
						StaminaMax:         8,
						ManaCurrent:        9,
						ManaMax:            10,
						Gold:               11,
						ResistancePhysical: 12,
						WeightCurrent:      13,
						Level3: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3{
							StatsCap: 14,
							Level4: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3_MsgStatWindowInfoLevel4{
								FollowersCurrent: 15,
								FollowersMax:     16,
								Level5: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3_MsgStatWindowInfoLevel4_MsgStatWindowInfoLevel5{
									ResistanceFire:   17,
									ResistanceCold:   18,
									ResistancePoison: 19,
									ResistanceEnergy: 20,
									Luck:             21,
									DamageMin:        22,
									DamageMax:        23,
									TithingPoints:    24,
									Level6: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3_MsgStatWindowInfoLevel4_MsgStatWindowInfoLevel5_MsgStatWindowInfoLevel6{
										WeightMax: 25,
										Race:      26,
										Level7:    nil,
									},
								},
							},
						},
					},
				},
			},
			want: commands.StatWindowInfo{
				CharacterID:                  42000,
				CharacterName:                "test",
				HitPointsCurrent:             1,
				HitPointsMax:                 2,
				FlagNameAllowed:              true,
				FlagDisplay:                  3,
				Gender:                       4,
				Strength:                     5,
				Intelligence:                 6,
				StaminaCurrent:               7,
				StaminaMax:                   8,
				ManaCurrent:                  9,
				ManaMax:                      10,
				Gold:                         11,
				ResistancePhysical:           12,
				WeightCurrent:                13,
				WeightMax:                    25,
				Race:                         26,
				StatsCap:                     14,
				FollowersCurrent:             15,
				FollowersMax:                 16,
				ResistanceFire:               17,
				ResistanceCold:               18,
				ResistancePoison:             19,
				ResistanceEnergy:             20,
				Luck:                         21,
				DamageMin:                    22,
				DamageMax:                    23,
				TithingPoints:                24,
				ResistancePhysicalMax:        0,
				ResistanceFireMax:            0,
				ResistanceColdMax:            0,
				ResistancePoisonMax:          0,
				ResistanceEnergyMax:          0,
				DefenseChanceIncreaseCurrent: 0,
				DefenseChanceIncreaseMax:     0,
				HitChanceIncrease:            0,
				SwingSpeedIncrease:           0,
				DamageIncrease:               0,
				LowerReagentCost:             0,
				SpellDamageIncrease:          0,
				FasterCastRecovery:           0,
				FasterCasting:                0,
				LowerManaCost:                0,
			},
		},
		{
			name: "level5 case",
			args: args{
				info: &services.MsgStatWindowInfo{
					CharacterID: &services.ObjectId{
						Value: 42000,
					},
					CharacterName:    "test",
					HitPointsCurrent: 1,
					HitPointsMax:     2,
					FlagNameAllowed:  true,
					FlagDisplay:      3,
					Level2: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2{
						Gender:             4,
						Strength:           5,
						Intelligence:       6,
						StaminaCurrent:     7,
						StaminaMax:         8,
						ManaCurrent:        9,
						ManaMax:            10,
						Gold:               11,
						ResistancePhysical: 12,
						WeightCurrent:      13,
						Level3: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3{
							StatsCap: 14,
							Level4: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3_MsgStatWindowInfoLevel4{
								FollowersCurrent: 15,
								FollowersMax:     16,
								Level5: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3_MsgStatWindowInfoLevel4_MsgStatWindowInfoLevel5{
									ResistanceFire:   17,
									ResistanceCold:   18,
									ResistancePoison: 19,
									ResistanceEnergy: 20,
									Luck:             21,
									DamageMin:        22,
									DamageMax:        23,
									TithingPoints:    24,
									Level6:           nil,
								},
							},
						},
					},
				},
			},
			want: commands.StatWindowInfo{
				CharacterID:                  42000,
				CharacterName:                "test",
				HitPointsCurrent:             1,
				HitPointsMax:                 2,
				FlagNameAllowed:              true,
				FlagDisplay:                  3,
				Gender:                       4,
				Strength:                     5,
				Intelligence:                 6,
				StaminaCurrent:               7,
				StaminaMax:                   8,
				ManaCurrent:                  9,
				ManaMax:                      10,
				Gold:                         11,
				ResistancePhysical:           12,
				WeightCurrent:                13,
				WeightMax:                    0,
				Race:                         0,
				StatsCap:                     14,
				FollowersCurrent:             15,
				FollowersMax:                 16,
				ResistanceFire:               17,
				ResistanceCold:               18,
				ResistancePoison:             19,
				ResistanceEnergy:             20,
				Luck:                         21,
				DamageMin:                    22,
				DamageMax:                    23,
				TithingPoints:                24,
				ResistancePhysicalMax:        0,
				ResistanceFireMax:            0,
				ResistanceColdMax:            0,
				ResistancePoisonMax:          0,
				ResistanceEnergyMax:          0,
				DefenseChanceIncreaseCurrent: 0,
				DefenseChanceIncreaseMax:     0,
				HitChanceIncrease:            0,
				SwingSpeedIncrease:           0,
				DamageIncrease:               0,
				LowerReagentCost:             0,
				SpellDamageIncrease:          0,
				FasterCastRecovery:           0,
				FasterCasting:                0,
				LowerManaCost:                0,
			},
		},
		{
			name: "level4 case",
			args: args{
				info: &services.MsgStatWindowInfo{
					CharacterID: &services.ObjectId{
						Value: 42000,
					},
					CharacterName:    "test",
					HitPointsCurrent: 1,
					HitPointsMax:     2,
					FlagNameAllowed:  true,
					FlagDisplay:      3,
					Level2: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2{
						Gender:             4,
						Strength:           5,
						Intelligence:       6,
						StaminaCurrent:     7,
						StaminaMax:         8,
						ManaCurrent:        9,
						ManaMax:            10,
						Gold:               11,
						ResistancePhysical: 12,
						WeightCurrent:      13,
						Level3: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3{
							StatsCap: 14,
							Level4: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3_MsgStatWindowInfoLevel4{
								FollowersCurrent: 15,
								FollowersMax:     16,
								Level5:           nil,
							},
						},
					},
				},
			},
			want: commands.StatWindowInfo{
				CharacterID:                  42000,
				CharacterName:                "test",
				HitPointsCurrent:             1,
				HitPointsMax:                 2,
				FlagNameAllowed:              true,
				FlagDisplay:                  3,
				Gender:                       4,
				Strength:                     5,
				Intelligence:                 6,
				StaminaCurrent:               7,
				StaminaMax:                   8,
				ManaCurrent:                  9,
				ManaMax:                      10,
				Gold:                         11,
				ResistancePhysical:           12,
				WeightCurrent:                13,
				WeightMax:                    0,
				Race:                         0,
				StatsCap:                     14,
				FollowersCurrent:             15,
				FollowersMax:                 16,
				ResistanceFire:               0,
				ResistanceCold:               0,
				ResistancePoison:             0,
				ResistanceEnergy:             0,
				Luck:                         0,
				DamageMin:                    0,
				DamageMax:                    0,
				TithingPoints:                0,
				ResistancePhysicalMax:        0,
				ResistanceFireMax:            0,
				ResistanceColdMax:            0,
				ResistancePoisonMax:          0,
				ResistanceEnergyMax:          0,
				DefenseChanceIncreaseCurrent: 0,
				DefenseChanceIncreaseMax:     0,
				HitChanceIncrease:            0,
				SwingSpeedIncrease:           0,
				DamageIncrease:               0,
				LowerReagentCost:             0,
				SpellDamageIncrease:          0,
				FasterCastRecovery:           0,
				FasterCasting:                0,
				LowerManaCost:                0,
			},
		},
		{
			name: "level3 case",
			args: args{
				info: &services.MsgStatWindowInfo{
					CharacterID: &services.ObjectId{
						Value: 42000,
					},
					CharacterName:    "test",
					HitPointsCurrent: 1,
					HitPointsMax:     2,
					FlagNameAllowed:  true,
					FlagDisplay:      3,
					Level2: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2{
						Gender:             4,
						Strength:           5,
						Intelligence:       6,
						StaminaCurrent:     7,
						StaminaMax:         8,
						ManaCurrent:        9,
						ManaMax:            10,
						Gold:               11,
						ResistancePhysical: 12,
						WeightCurrent:      13,
						Level3: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2_MsgStatWindowInfoLevel3{
							StatsCap: 14,
							Level4:   nil,
						},
					},
				},
			},
			want: commands.StatWindowInfo{
				CharacterID:                  42000,
				CharacterName:                "test",
				HitPointsCurrent:             1,
				HitPointsMax:                 2,
				FlagNameAllowed:              true,
				FlagDisplay:                  3,
				Gender:                       4,
				Strength:                     5,
				Intelligence:                 6,
				StaminaCurrent:               7,
				StaminaMax:                   8,
				ManaCurrent:                  9,
				ManaMax:                      10,
				Gold:                         11,
				ResistancePhysical:           12,
				WeightCurrent:                13,
				WeightMax:                    0,
				Race:                         0,
				StatsCap:                     14,
				FollowersCurrent:             0,
				FollowersMax:                 0,
				ResistanceFire:               0,
				ResistanceCold:               0,
				ResistancePoison:             0,
				ResistanceEnergy:             0,
				Luck:                         0,
				DamageMin:                    0,
				DamageMax:                    0,
				TithingPoints:                0,
				ResistancePhysicalMax:        0,
				ResistanceFireMax:            0,
				ResistanceColdMax:            0,
				ResistancePoisonMax:          0,
				ResistanceEnergyMax:          0,
				DefenseChanceIncreaseCurrent: 0,
				DefenseChanceIncreaseMax:     0,
				HitChanceIncrease:            0,
				SwingSpeedIncrease:           0,
				DamageIncrease:               0,
				LowerReagentCost:             0,
				SpellDamageIncrease:          0,
				FasterCastRecovery:           0,
				FasterCasting:                0,
				LowerManaCost:                0,
			},
		},
		{
			name: "level2 case",
			args: args{
				info: &services.MsgStatWindowInfo{
					CharacterID: &services.ObjectId{
						Value: 42000,
					},
					CharacterName:    "test",
					HitPointsCurrent: 1,
					HitPointsMax:     2,
					FlagNameAllowed:  true,
					FlagDisplay:      3,
					Level2: &services.MsgStatWindowInfo_MsgStatWindowInfoLevel2{
						Gender:             4,
						Strength:           5,
						Intelligence:       6,
						StaminaCurrent:     7,
						StaminaMax:         8,
						ManaCurrent:        9,
						ManaMax:            10,
						Gold:               11,
						ResistancePhysical: 12,
						WeightCurrent:      13,
						Level3:             nil,
					},
				},
			},
			want: commands.StatWindowInfo{
				CharacterID:                  42000,
				CharacterName:                "test",
				HitPointsCurrent:             1,
				HitPointsMax:                 2,
				FlagNameAllowed:              true,
				FlagDisplay:                  3,
				Gender:                       4,
				Strength:                     5,
				Intelligence:                 6,
				StaminaCurrent:               7,
				StaminaMax:                   8,
				ManaCurrent:                  9,
				ManaMax:                      10,
				Gold:                         11,
				ResistancePhysical:           12,
				WeightCurrent:                13,
				WeightMax:                    0,
				Race:                         0,
				StatsCap:                     0,
				FollowersCurrent:             0,
				FollowersMax:                 0,
				ResistanceFire:               0,
				ResistanceCold:               0,
				ResistancePoison:             0,
				ResistanceEnergy:             0,
				Luck:                         0,
				DamageMin:                    0,
				DamageMax:                    0,
				TithingPoints:                0,
				ResistancePhysicalMax:        0,
				ResistanceFireMax:            0,
				ResistanceColdMax:            0,
				ResistancePoisonMax:          0,
				ResistanceEnergyMax:          0,
				DefenseChanceIncreaseCurrent: 0,
				DefenseChanceIncreaseMax:     0,
				HitChanceIncrease:            0,
				SwingSpeedIncrease:           0,
				DamageIncrease:               0,
				LowerReagentCost:             0,
				SpellDamageIncrease:          0,
				FasterCastRecovery:           0,
				FasterCasting:                0,
				LowerManaCost:                0,
			},
		},
		{
			name: "level1 case",
			args: args{
				info: &services.MsgStatWindowInfo{
					CharacterID: &services.ObjectId{
						Value: 42000,
					},
					CharacterName:    "test",
					HitPointsCurrent: 1,
					HitPointsMax:     2,
					FlagNameAllowed:  true,
					FlagDisplay:      3,
					Level2:           nil,
				},
			},
			want: commands.StatWindowInfo{
				CharacterID:                  42000,
				CharacterName:                "test",
				HitPointsCurrent:             1,
				HitPointsMax:                 2,
				FlagNameAllowed:              true,
				FlagDisplay:                  3,
				Gender:                       0,
				Strength:                     0,
				Intelligence:                 0,
				StaminaCurrent:               0,
				StaminaMax:                   0,
				ManaCurrent:                  0,
				ManaMax:                      0,
				Gold:                         0,
				ResistancePhysical:           0,
				WeightCurrent:                0,
				WeightMax:                    0,
				Race:                         0,
				StatsCap:                     0,
				FollowersCurrent:             0,
				FollowersMax:                 0,
				ResistanceFire:               0,
				ResistanceCold:               0,
				ResistancePoison:             0,
				ResistanceEnergy:             0,
				Luck:                         0,
				DamageMin:                    0,
				DamageMax:                    0,
				TithingPoints:                0,
				ResistancePhysicalMax:        0,
				ResistanceFireMax:            0,
				ResistanceColdMax:            0,
				ResistancePoisonMax:          0,
				ResistanceEnergyMax:          0,
				DefenseChanceIncreaseCurrent: 0,
				DefenseChanceIncreaseMax:     0,
				HitChanceIncrease:            0,
				SwingSpeedIncrease:           0,
				DamageIncrease:               0,
				LowerReagentCost:             0,
				SpellDamageIncrease:          0,
				FasterCastRecovery:           0,
				FasterCasting:                0,
				LowerManaCost:                0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapToCommandStatWindowInfo(tt.args.info); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapToCommandStatWindowInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
