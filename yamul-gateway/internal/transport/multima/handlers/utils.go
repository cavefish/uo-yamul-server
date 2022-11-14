package handlers

import (
	"strconv"
	"strings"
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
