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
	Now *now.Now
}

func (calendar *Calendar) Next() {
	newDate := calendar.Now.BeginningOfMonth().AddDate(0, 1, 0)
	calendar.Now = now.New(newDate)
}

func (calendar *Calendar) Previous() {
	newDate := calendar.Now.BeginningOfMonth().AddDate(0, -1, 0)
	calendar.Now = now.New(newDate)
}

func (calendar *Calendar) Week() (week Week) {
	beginningOfWeek := calendar.Now.BeginningOfWeek()

	for i := 0; i < 7; i++ {
		week = append(week, beginningOfWeek)
		beginningOfWeek = beginningOfWeek.AddDate(0, 0, 1)
	}
	return
}

func (calendar *Calendar) NextWeek() (week Week) {
	newDate := calendar.Now.AddDate(0, 0, 7)
	calendar.Now = now.New(newDate)
	defer func() {
		calendar.Now = now.New(newDate.AddDate(0, 0, -7))
	}()

	week = calendar.Week()
	return
}

func (calendar *Calendar) PreviousWeek() (week Week) {
	newDate := calendar.Now.AddDate(0, 0, -7)
	calendar.Now = now.New(newDate)
	defer func() {
		calendar.Now = now.New(newDate.AddDate(0, 0, 7))
	}()

	week = calendar.Week()
	return
}

func (calendar *Calendar) Month() (month Month) {
	beginningOfMonth := calendar.Now.BeginningOfMonth()
	endOfMonth := calendar.Now.EndOfMonth()
	week := New(beginningOfMonth).Week()
	lastWeek := New(endOfMonth).Week()

	for !reflect.DeepEqual(lastWeek, week) {
		month = append(month, week)
		week = week.Next()
	}
	month = append(month, lastWeek)
	return
}

func (calendar *Calendar) NextMonth() (month Month) {
	calendar.Next()
	defer calendar.Previous()
	month = calendar.Month()
	return
}

func (calendar *Calendar) PreviousMonth() (month Month) {
	calendar.Previous()
	defer calendar.Next()
	month = calendar.Month()
	return
}

func New(t time.Time) *Calendar {
	return &Calendar{
		Now: now.New(t),
	}
}
