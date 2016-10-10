package calendar

import (
	"reflect"
	"time"

	"github.com/jinzhu/now"
)

// Week have time.Time data to represent week.
type Week []time.Time

// Next returns time.Time collection to represent next week.
func (week Week) Next() (nextWeek Week) {
	for _, t := range week {
		nextWeek = append(nextWeek, t.AddDate(0, 0, 7))
	}
	return
}

// Previous returns time.Time collection to represent previous week.
func (week Week) Previous() (previousWeek Week) {
	for _, t := range week {
		previousWeek = append(previousWeek, t.AddDate(0, 0, -7))
	}
	return
}

// Month have Week data to represent month.
type Month []Week

// Calendar have now.Now data.
type Calendar struct {
	Now *now.Now
}

// Next sets new *now.Now for next month.
func (calendar *Calendar) Next() {
	newDate := calendar.Now.BeginningOfMonth().AddDate(0, 1, 0)
	calendar.Now = now.New(newDate)
}

// Previous sets new *now.Now for previous month.
func (calendar *Calendar) Previous() {
	newDate := calendar.Now.BeginningOfMonth().AddDate(0, -1, 0)
	calendar.Now = now.New(newDate)
}

// NextCalendar returns next Calendar.
func (calendar *Calendar) NextCalendar() (nextCalendar *Calendar) {
	newDate := calendar.Now.BeginningOfMonth().AddDate(0, 1, 0)
	nextCalendar = &Calendar{
		Now: now.New(newDate),
	}
	return
}

// PreviousCalendar returns previous Calendar.
func (calendar *Calendar) PreviousCalendar() (previousCalendar *Calendar) {
	newDate := calendar.Now.BeginningOfMonth().AddDate(0, -1, 0)
	previousCalendar = &Calendar{
		Now: now.New(newDate),
	}
	return
}

// Week returns Week regarding current date.
func (calendar *Calendar) Week() (week Week) {
	beginningOfWeek := calendar.Now.BeginningOfWeek()

	for i := 0; i < 7; i++ {
		week = append(week, beginningOfWeek)
		beginningOfWeek = beginningOfWeek.AddDate(0, 0, 1)
	}
	return
}

// NextWeek returns next Week regarding current date.
// It doesn't have side effect.
func (calendar *Calendar) NextWeek() (week Week) {
	newDate := calendar.Now.AddDate(0, 0, 7)
	calendar.Now = now.New(newDate)
	defer func() {
		calendar.Now = now.New(newDate.AddDate(0, 0, -7))
	}()

	week = calendar.Week()
	return
}

// PreviousWeek returns previous Week regarding current date.
// It doesn't have side effect.
func (calendar *Calendar) PreviousWeek() (week Week) {
	newDate := calendar.Now.AddDate(0, 0, -7)
	calendar.Now = now.New(newDate)
	defer func() {
		calendar.Now = now.New(newDate.AddDate(0, 0, 7))
	}()

	week = calendar.Week()
	return
}

// Month returns Month regarding current date.
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

// NextMonth returns next Month regarding current date.
func (calendar *Calendar) NextMonth() (month Month) {
	month = calendar.NextCalendar().Month()
	return
}

// PreviousMonth returns previous Month regarding current date.
func (calendar *Calendar) PreviousMonth() (month Month) {
	month = calendar.PreviousCalendar().Month()
	return
}
