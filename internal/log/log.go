package log

import (
	"fmt"
	"os"
)

var enabled = os.Getenv("XRM_DEBUG") == "1"

func Info(format string, args ...any) {
	if enabled {
		fmt.Printf("[XRM] "+format+"\n", args...)
	}
}
