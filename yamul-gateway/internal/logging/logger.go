package logging

import "fmt"

func Info(format string, vars ...any) {
	fmt.Printf(format, vars...)
}

func Debug(format string, vars ...any) {
	fmt.Printf(format, vars...)
}
