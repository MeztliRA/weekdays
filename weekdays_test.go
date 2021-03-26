package weekdays_test

import (
	"testing"
	"time"

	"github.com/MeztliRA/weekdays"
)

const (
	messageWeekday = "Its the weekday!"
	messageWeekend = "Its the weekend!"
)

func TestMessage(t *testing.T) {
	got := weekdays.Message()
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		if got != messageWeekend {
			t.Errorf("Message is not the same")
		}
	default:
		if got != messageWeekday {
			t.Errorf("Message is not the same")
		}
	}
}
