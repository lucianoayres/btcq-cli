package utils

import (
	"fmt"
	"time"
)

func LoadingAnimation(stopChan chan bool) {
	// Define characters for the animation
	frames := []string{
		"\r[   ]",
		"\r[.  ]",
		"\r[.. ]",
		"\r[...]",
	}
	i := 0
	for {
		select {
		case <-stopChan:
			// Clear the spinner from the line
			fmt.Print("\r          \r") // Clear the line
			return
		default:
			// Print the current animation frame
			fmt.Print(frames[i])
			i = (i + 1) % len(frames)
			time.Sleep(200 * time.Millisecond)
		}
	}
}