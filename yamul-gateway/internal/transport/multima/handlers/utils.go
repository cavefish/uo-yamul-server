package handlers

import (
	"strconv"
	"strings"
	"yamul-gateway/internal/dtos/commands"
)

/*
*
Returns Address in LE, and port number in LE
*/
func addressToUInt(value string) (uint32, uint16) {
	ip := strings.Split(value, ":")
	ipTokens := strings.Split(ip[0], ".")
	var result uint32 = 0
	for i := 0; i < 4; i++ {
		v, _ := strconv.Atoi(ipTokens[i])
		result = result<<8 | uint32(v)
	}
	port, _ := strconv.Atoi(ip[1])
	return result, uint16(port)
}

func ConvertClientFeaturesToFlags(features commands.ClientFeatures) uint32 {
	var flags uint32 = 0
	flags |= ifFlagIsSet(0x0001, features.Unknown0001)
	flags |= ifFlagIsSet(0x0002, features.OverwriteConfigButtons)
	flags |= ifFlagIsSet(0x0004, features.SingleCharacterSlot)
	flags |= ifFlagIsSet(0x0008, features.ContextMenus)
	flags |= ifFlagIsSet(0x0010, features.LimitCharacterSlots)
	flags |= ifFlagIsSet(0x0020, features.EnableAOS)
	flags |= ifFlagIsSet(0x0040, features.SixthSlot)
	flags |= ifFlagIsSet(0x0080, features.SamuraiNinja)
	flags |= ifFlagIsSet(0x0100, features.ElvenRace)
	flags |= ifFlagIsSet(0x0200, features.Unknown0200)
	flags |= ifFlagIsSet(0x0400, features.UO3D)
	flags |= ifFlagIsSet(0x0800, features.Unknown0800)
	flags |= ifFlagIsSet(0x1000, features.SeventhSlot)
	flags |= ifFlagIsSet(0x2000, features.Unknown2000)
	flags |= ifFlagIsSet(0x4000, features.NewMovement)
	flags |= ifFlagIsSet(0x8000, features.UnlockPvPAreas)
	return flags
}

func ifFlagIsSet(mask uint32, flag bool) uint32 {
	if flag {
		return mask
	}
	return 0
}
