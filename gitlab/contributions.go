package gitlab

import (
	"time"

	"github.com/xanzy/go-gitlab"
)

func (g *Gitlab) GetContributions(date time.Time) ([]*gitlab.ContributionEvent, error) {
	if err := g.getClient(); err != nil {
		return nil, err
	}

	start := gitlab.ISOTime(date.Add(time.Duration(-24 * time.Hour)))
	end := gitlab.ISOTime(date.Add(time.Duration(24 * time.Hour)))

	events, _, err := g.client.Events.ListCurrentUserContributionEvents(&gitlab.ListContributionEventsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 2500,
		},
		After:  &start,
		Before: &end,
	})
	if err != nil {
		return nil, err
	}

	return events, nil
}
