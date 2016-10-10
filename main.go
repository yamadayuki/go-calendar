package calendar

import (
	"time"

	"github.com/jinzhu/now"
)

// New returns new Calendar pointer.
func New(t time.Time) *Calendar {
	return &Calendar{
		Now: now.New(t),
	}
}

// Now returns Calendar
func Now() *Calendar {
	return &Calendar{
		Now: now.New(time.Now()),
	}
}

// CurrentMonth returns Now().Month()
func CurrentMonth() Month {
	return Now().Month()
}

// NextMonth returns Now().NextMonth()
func NextMonth() Month {
	return Now().NextMonth()
}

// PreviousMonth returns Now().PreviousMonth()
func PreviousMonth() Month {
	return Now().PreviousMonth()
}

// CurrentWeek returns Now().Week()
func CurrentWeek() Week {
	return Now().Week()
}

// NextWeek returns Now().NextWeek()
func NextWeek() Week {
	return Now().NextWeek()
}

// PreviousWeek returns Now().PreviousWeek()
func PreviousWeek() Week {
	return Now().PreviousWeek()
}
