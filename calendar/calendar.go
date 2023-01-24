package calendar

import (
	"github.com/fouad1998/work-reporter/config"
	cal "google.golang.org/api/calendar/v3"
)

type Calendar struct {
	service *cal.Service
	config  config.Config
}
