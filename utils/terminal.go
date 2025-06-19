package utils

import "fmt"

func ClearScreen() {
	fmt.Print("\x1b[2J\x1b[0f")
}
