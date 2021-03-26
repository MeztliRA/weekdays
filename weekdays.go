package weekdays

import (
	"time"
)

const (
	messageWeekday = "Its the weekday!"
	messageWeekend = "Its the weekend!"
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
		return messageWeekend
	default:
		return messageWeekday
	}
}
