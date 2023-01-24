package jira

import (
	"fmt"
	"time"

	J "github.com/andygrunwald/go-jira"
)

func (j *Jira) GetTasks(date time.Time) ([]J.Issue, error) {
	if j.client == nil {
		if err := j.getClient(); err != nil {
			return nil, err
		}
	}

	if err := j.config.Read(); err != nil {
		return nil, err
	}

	issues, _, err := j.client.Issue.Search(fmt.Sprintf(` 
		project = "CUR" AND assignee = currentUser()
		AND statusCategory in ("In Progress","Complete")
		AND type = "Subtask" AND status IN ("Done","In Progress")
		AND updated >= "%s" AND updated < "%s"  ORDER BY created DESC
		`, date.Format("2006-01-02"), date.Add(24*time.Hour).Format("2006-01-02")),
		&J.SearchOptions{MaxResults: 1000})

	if err != nil {
		return nil, err
	}

	return issues, nil
}
