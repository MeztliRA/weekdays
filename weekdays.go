package weekdays

import (
	"time"
)

func IsWeekday() bool {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		return false
	default:
		return true
	}
}

func IsWeekend() bool {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		return true
	default:
		return false
	}
}

func Message() string {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		return "Its the weekend!"
	default:
		return "Its the weekday!"
	}
}
