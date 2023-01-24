package calendar

import (
	"time"

	"google.golang.org/api/calendar/v3"
)

func (c *Calendar) GetTodayEvents() ([]*calendar.Event, error) {
	return c.GetEvents(time.Now())
}
