package utils

import (
	"testing"
	"time"
)

func TestLoadingAnimationCompletion(t *testing.T) {
	// Create a channel to stop the animation
	stopChan := make(chan bool)

	// Start the animation in a goroutine
	go func() {
		LoadingAnimation(stopChan)
	}()

	// Allow the animation to run for a short time
	time.Sleep(600 * time.Millisecond)

	// Stop the animation
	stopChan <- true

	// Allow some time for the animation to complete
	time.Sleep(100 * time.Millisecond)

	// The test passes if we reach this point without any issues
	// No specific output check is needed
}
