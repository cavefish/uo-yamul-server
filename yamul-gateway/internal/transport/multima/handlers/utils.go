package handlers

import (
	"strconv"
	"strings"
	"yamul-gateway/internal/transport/multima/commands"
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

func convertClientFeaturesToFlags(features commands.ClientFeatures) uint32 {
	var flags uint32 = 0
	if features.Chat {
		flags |= 0x0001
	}
	if features.LbrAnimations {
		flags |= 0x0002
	}
	if features.CreatePaladinNecromancer {
		flags |= 0x0010
	}
	if features.SixthSlot {
		flags |= 0x0020
	}
	if features.ExtraFeatures {
		flags |= 0x8000
	}
	return flags
}
