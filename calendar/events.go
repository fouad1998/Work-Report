package calendar

import (
	"time"

	"google.golang.org/api/calendar/v3"
)

func (c *Calendar) GetEvents(date time.Time) ([]*calendar.Event, error) {
	if c.service == nil {
		if err := c.getClient(); err != nil {
			return nil, err
		}
	}

	year, month, day := date.Date()

	start := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	end := time.Date(year, month, day, 23, 59, 59, 0, time.Local)

	events, err := c.service.Events.List("primary").
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(start.Format(time.RFC3339)).
		TimeMax(end.Format(time.RFC3339)).
		MaxResults(100).
		OrderBy("startTime").Do()
	if err != nil {
		return nil, err
	}

	return events.Items, nil
}
