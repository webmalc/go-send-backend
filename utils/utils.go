package utils

import (
	"fmt"
)

// ProcessFatalError checks  fatal errors
func ProcessFatalError(err error) {
	if err != nil {
		panic(fmt.Errorf("error: %s", err))
	}
}
