package booleans

func BoolToByte(b bool) byte {
	if b {
		return 0xff
	}
	return 0
}
