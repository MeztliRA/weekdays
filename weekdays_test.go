package weekdays_test

import (
	"testing"
	"time"

	"github.com/MeztliRA/weekdays"
)

const (
	messageWeekday      = "Its the weekday!"
	messageWeekend      = "Its the weekend!"
	messageWeekdayShort = "weekday!"
	messageWeekendShort = "weekend!"
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

func TestMessageShort(t *testing.T) {
	got := weekdays.MessageShort()
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		if got != messageWeekendShort {
			t.Errorf("Message is not the same")
		}
	default:
		if got != messageWeekdayShort {
			t.Errorf("Message is not the same")
		}
	}
}
