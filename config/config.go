package config

import "golang.org/x/oauth2"

type Config struct {
	JiraToken   string        `json:"jira_token"`
	JiraEmail   string        `json:"jira_email"`
	GitlabToken string        `json:"g_lab_token"`
	GToken      *oauth2.Token `json:"g_token"`
}
