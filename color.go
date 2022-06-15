package logf

import "fmt"

type COLOR uint8

const (
	COLOR_RED     COLOR = 91
	COLOR_GREEN   COLOR = 92
	COLOR_YELLOW  COLOR = 93
	COLOR_BLUE    COLOR = 94
	COLOR_MAGENTA COLOR = 95
)

// debug
func blue(msg string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", COLOR_BLUE, msg)
}

// info
func green(msg string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", COLOR_GREEN, msg)
}

// error
func red(msg string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", COLOR_RED, msg)
}

// warn
func yellow(msg string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", COLOR_YELLOW, msg)
}
