package weekdays

import (
	"fmt"
	"time"
)

const (
	messageWeekday      = "Its the weekday!"
	messageWeekend      = "Its the weekend!"
	messageWeekdayShort = "weekday!"
	messageWeekendShort = "weekend!"
)

// return true if its the weekday, else return false
func IsWeekday() bool {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		return false
	default:
		return true
	}
}

// return true if its the weekend, else return false
func IsWeekend() bool {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		return true
	default:
		return false
	}
}

// return a different message depending on whether its the weekday or the weekend
func Message() string {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		return messageWeekend
	default:
		return messageWeekday
	}
}

// similar to Message(), but the returned message are shorter
func MessageShort() string {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		return messageWeekendShort
	default:
		return messageWeekdayShort
	}
}

// print the returned message from Message()
func PrintMessage() {
	fmt.Println(Message())
}

// print the returned message from MessageShort()
func PrintMessageShort() {
	fmt.Println(MessageShort())
}
