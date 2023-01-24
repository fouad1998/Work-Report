package reporter

import (
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/xanzy/go-gitlab"
	"google.golang.org/api/calendar/v3"
)

type Report struct {
	Date          time.Time
	Issues        []jira.Issue
	Events        []*calendar.Event
	Contributions []*gitlab.ContributionEvent
	Note          string
	Hours         int
}
