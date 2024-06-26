package stringUtils

func TrimRight(str string) string {
	var i = len(str) - 1
	for {
		if i < 0 {
			return ""
		}
		if str[i] != 0 {
			return str[0 : i+1]
		}
		i--
	}
}

func RightPaddingByte(str string, length int) []byte {
	var result = make([]byte, length)
	var i = 0
	for i < length && i < len(str) {
		result[i] = str[i]
		i++
	}
	for i < length {
		result[i] = 0
		i++
	}
	return result
}
