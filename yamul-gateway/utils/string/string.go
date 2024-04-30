package string

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
