package jira

import (
	J "github.com/andygrunwald/go-jira"
)

func (j *Jira) getClient() error {
	if err := j.config.Read(); err != nil {
		return err
	}

	tp := J.BasicAuthTransport{
		Username: j.config.JiraEmail,
		Password: j.config.JiraToken,
	}

	client, err := J.NewClient(tp.Client(), "https://cureety.atlassian.net")
	if err != nil {
		return err
	}

	j.client = client
	return nil
}
