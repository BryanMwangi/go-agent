package utils

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

var (
	TermSpinner *spinner.Spinner
)

func ShowLoader(message string) {
	fmt.Println(message)
	TermSpinner = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	TermSpinner.Start()
}

// This is used to stop the loader after a short delay
//
// Please note that this is not good practice because if the TermSpinner is not
// initialized, it will cause a panic
//
// I just did this for the sake of simplicity and to reduce overall memory usage
func StopLoader(sleep time.Duration) {
	time.Sleep(sleep)
	TermSpinner.Stop()
}

func ClearScreen() {
	fmt.Print("\x1b[2J\x1b[0f")
}
