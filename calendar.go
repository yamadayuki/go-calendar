package calendar

import (
	"reflect"
	"time"

	"github.com/jinzhu/now"
)

type Week []time.Time

func (week Week) Next() (nextWeek Week) {
	for _, t := range week {
		nextWeek = append(nextWeek, t.AddDate(0, 0, 7))
	}
	return
}

func (week Week) Previous() (previousWeek Week) {
	for _, t := range week {
		previousWeek = append(previousWeek, t.AddDate(0, 0, -7))
	}
	return
}

type Month []Week

type Calendar struct {
	Year        int
	MonthNumber time.Month
	Now         *now.Now
}

func (calendar *Calendar) Week() (week Week) {
	beginningOfWeek := calendar.Now.BeginningOfWeek()

	for i := 0; i < 7; i++ {
		week = append(week, beginningOfWeek)
		beginningOfWeek = beginningOfWeek.AddDate(0, 0, 1)
	}
	return
}

func (calendar *Calendar) Month() (weeks Month) {
	beginningOfMonth := calendar.Now.BeginningOfMonth()
	endOfMonth := calendar.Now.EndOfMonth()
	week := New(beginningOfMonth).Week()
	lastWeek := New(endOfMonth).Week()

	for !reflect.DeepEqual(lastWeek, week) {
		weeks = append(weeks, week)
		week = week.Next()
	}
	weeks = append(weeks, lastWeek)
	return
}

func New(t time.Time) *Calendar {
	return &Calendar{
		Year:        t.Year(),
		MonthNumber: t.Month(),
		Now:         now.New(t),
	}
}
