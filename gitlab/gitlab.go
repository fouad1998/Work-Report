package gitlab

import (
	"github.com/fouad1998/work-reporter/config"
	"github.com/xanzy/go-gitlab"
)

type Gitlab struct {
	config config.Config
	client *gitlab.Client
}
