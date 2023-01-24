package gitlab

import (
	"github.com/xanzy/go-gitlab"
)

func (g *Gitlab) getClient() error {
	if err := g.config.Read(); err != nil {
		return err
	}

	git, err := gitlab.NewClient(g.config.GitlabToken)
	if err != nil {
		return err
	}

	g.client = git

	return nil
}
