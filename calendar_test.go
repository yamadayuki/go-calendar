package calendar

import (
	"reflect"
	"testing"
	"time"
)

var date = time.Date(2016, 10, 7, 17, 51, 49, 123456789, time.UTC)

var expectedWeekDays = Week{
	time.Date(2016, 10, 2, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 3, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 5, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 6, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 7, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 8, 0, 0, 0, 0, time.UTC),
}

var expectedMonth = Month{
	Week{
		time.Date(2016, 9, 25, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 9, 26, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 9, 27, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 9, 28, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 9, 29, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 9, 30, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 1, 0, 0, 0, 0, time.UTC),
	},
	Week{
		time.Date(2016, 10, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 7, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 8, 0, 0, 0, 0, time.UTC),
	},
	Week{
		time.Date(2016, 10, 9, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 10, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 11, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 12, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 13, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 14, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 15, 0, 0, 0, 0, time.UTC),
	},
	Week{
		time.Date(2016, 10, 16, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 17, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 18, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 19, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 20, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 21, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 22, 0, 0, 0, 0, time.UTC),
	},
	Week{
		time.Date(2016, 10, 23, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 24, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 25, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 26, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 27, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 28, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 29, 0, 0, 0, 0, time.UTC),
	},
	Week{
		time.Date(2016, 10, 30, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 5, 0, 0, 0, 0, time.UTC),
	},
}

var expectedNextMonth = Month{
	Week{
		time.Date(2016, 10, 30, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 10, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 5, 0, 0, 0, 0, time.UTC),
	},
	Week{
		time.Date(2016, 11, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 7, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 8, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 9, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 10, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 11, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 12, 0, 0, 0, 0, time.UTC),
	},
	Week{
		time.Date(2016, 11, 13, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 14, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 15, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 16, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 17, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 18, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 19, 0, 0, 0, 0, time.UTC),
	},
	Week{
		time.Date(2016, 11, 20, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 21, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 22, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 23, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 24, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 25, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 26, 0, 0, 0, 0, time.UTC),
	},
	Week{
		time.Date(2016, 11, 27, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 28, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 29, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 11, 30, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 12, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 12, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2016, 12, 3, 0, 0, 0, 0, time.UTC),
	},
}

var expectedNextWeekDays = Week{
	time.Date(2016, 10, 9, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 10, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 11, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 12, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 13, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 14, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 15, 0, 0, 0, 0, time.UTC),
}

var expectedPreviousWeekDays = Week{
	time.Date(2016, 9, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 9, 26, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 9, 27, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 9, 28, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 9, 29, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 9, 30, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 1, 0, 0, 0, 0, time.UTC),
}

func TestNext(t *testing.T) {
	cal := New(date)

	month := int(cal.Now.Month())
	if month != 10 {
		t.Errorf("Unexpected month, \n expected %v, \n given %v", 10, month)
	}

	cal.Next()

	nextMonth := int(cal.Now.Month())
	if nextMonth != 11 {
		t.Errorf("Unexpected Calendar.Next(), \n expected %v, \n given %v", 11, nextMonth)
	}
}

func TestPrevious(t *testing.T) {
	cal := New(date)

	month := int(cal.Now.Month())
	if month != 10 {
		t.Errorf("Unexpected month, \n expected %v, \n given %v", 10, month)
	}

	cal.Previous()

	previousMonth := int(cal.Now.Month())
	if previousMonth != 9 {
		t.Errorf("Unexpected Calendar.Previous(), \n expected %v, \n given %v", 9, previousMonth)
	}
}

func TestNextWeek(t *testing.T) {
	cal := New(date)

	nextWeek := cal.NextWeek()
	if !reflect.DeepEqual(expectedNextWeekDays, nextWeek) {
		t.Errorf("Unexpected Calendar.NextWeek(), \n expected %v, \n given %v", expectedNextWeekDays, nextWeek)
	}

	nextWeek = cal.Week().Next()
	if !reflect.DeepEqual(expectedNextWeekDays, nextWeek) {
		t.Errorf("Unexpected Calendar.Week().Next(), \n expected %v, \n given %v", expectedNextWeekDays, nextWeek)
	}

	week := cal.Week()
	if !reflect.DeepEqual(expectedWeekDays, week) {
		t.Errorf("Unexpected Calendar.Week(), \n expected %v, \n given %v", expectedWeekDays, week)
	}
}

func TestPreviousWeek(t *testing.T) {
	cal := New(date)

	previousWeek := cal.PreviousWeek()
	if !reflect.DeepEqual(expectedPreviousWeekDays, previousWeek) {
		t.Errorf("Unexpected Calendar.PreviousWeek(), \n expected %v, \n given %v", expectedPreviousWeekDays, previousWeek)
	}

	previousWeek = cal.Week().Previous()
	if !reflect.DeepEqual(expectedPreviousWeekDays, previousWeek) {
		t.Errorf("Unexpected Calendar.Week().Previous(), \n expected %v, \n given %v", expectedPreviousWeekDays, previousWeek)
	}

	week := cal.Week()
	if !reflect.DeepEqual(expectedWeekDays, week) {
		t.Errorf("Unexpected Calendar.Week(), \n expected %v, \n given %v", expectedWeekDays, week)
	}
}

func TestWeek(t *testing.T) {
	week := New(date).Week()
	if !reflect.DeepEqual(expectedWeekDays, week) {
		t.Errorf("Unexpected Calendar.Week(), \n expected %v, \n given %v", expectedWeekDays, week)
	}
}

func TestMonth(t *testing.T) {
	month := New(date).Month()
	if !reflect.DeepEqual(expectedMonth, month) {
		t.Errorf("Unexpected Calendar.Month(), \n expected %v, \n given %v", expectedMonth, month)
	}
}
