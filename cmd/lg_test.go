package cmd

import "testing"

func TestLogger(t *testing.T) {
	InitLogger(true)
	log.Debug("Debug message.")
	InitLogger(false)
	log.Debug("Debug message.")
}
