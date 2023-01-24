package gitlab

import (
	"time"

	"github.com/xanzy/go-gitlab"
)

func (g *Gitlab) GetTodayContribution() ([]*gitlab.ContributionEvent, error) {
	return g.GetContributions(time.Now())
}
