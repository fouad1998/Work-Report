package jira

import (
	jiira "github.com/andygrunwald/go-jira"
	"github.com/fouad1998/work-reporter/config"
)

type Jira struct {
	client *jiira.Client
	config config.Config
}
