package calendar

import (
	"reflect"
	"testing"
	"time"
)

var expectedWeekDays = Week{
	time.Date(2016, 10, 2, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 3, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 5, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 6, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 7, 0, 0, 0, 0, time.UTC),
	time.Date(2016, 10, 8, 0, 0, 0, 0, time.UTC),
}

var expectedWeeks = Weeks{
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

func TestNextWeek(t *testing.T) {
	n := time.Date(2016, 10, 7, 17, 51, 49, 123456789, time.UTC)

	nextWeek := New(n).Week().Next()
	if !reflect.DeepEqual(expectedNextWeekDays, nextWeek) {
		t.Errorf("Unexpected Week.Next(), \n expected %v, \n given %v", expectedNextWeekDays, nextWeek)
	}
}

func TestPreviousWeek(t *testing.T) {
	n := time.Date(2016, 10, 7, 17, 51, 49, 123456789, time.UTC)

	previousWeek := New(n).Week().Previous()
	if !reflect.DeepEqual(expectedPreviousWeekDays, previousWeek) {
		t.Errorf("Unexpected Week.Previous(), \n expected %v, \n given %v", expectedPreviousWeekDays, previousWeek)
	}
}

func TestWeek(t *testing.T) {
	n := time.Date(2016, 10, 7, 17, 51, 49, 123456789, time.UTC)

	week := New(n).Week()
	if !reflect.DeepEqual(expectedWeekDays, week) {
		t.Errorf("Unexpected Calendar.Week(), \n expected %v, \n given %v", expectedWeekDays, week)
	}
}

func TestWeeks(t *testing.T) {
	n := time.Date(2016, 10, 7, 17, 51, 49, 123456789, time.UTC)

	weeks := New(n).Weeks()
	if !reflect.DeepEqual(expectedWeeks, weeks) {
		t.Errorf("Unexpected Calendar.Weeks(), \n expected %v, \n given %v", expectedWeeks, weeks)
	}
}
