package jira

import (
	"time"

	J "github.com/andygrunwald/go-jira"
)

func (j *Jira) GetTodayTasks() ([]J.Issue, error) {
	return j.GetTasks(time.Now())
}
