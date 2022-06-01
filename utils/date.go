package utils

import (
	"fmt"
	"time"
)

func GetFormatTime(format string, time time.Time) string {
	return fmt.Sprintf(format, time.Year(), int(time.Month()), time.Day())
}
