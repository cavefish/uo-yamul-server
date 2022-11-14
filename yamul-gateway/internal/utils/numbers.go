package utils

func BigEndianInt(n int) int {
	return ((n & 0xFF) << 24) | ((n & 0xFF00) << 8) | ((n & 0xFF0000) >> 8) | (n >> 24)
}

func BigEndianUInt32(n uint32) uint32 {
	return ((n & 0xFF) << 24) | ((n & 0xFF00) << 8) | ((n & 0xFF0000) >> 8) | (n >> 24)
}
