package reporter

import (
	"strings"

	"github.com/xanzy/go-gitlab"
)

func strContr(c *gitlab.ContributionEvent) string {
	if c.TargetTitle != "" {
		return c.TargetTitle
	}

	if c.Title != "" {
		return c.Title
	}

	if c.PushData.CommitTitle != "" {
		return c.PushData.CommitTitle
	}

	if c.PushData.CommitTo != "" {
		return c.PushData.CommitTo
	}

	return c.PushData.Ref
}

func (r *Report) object() Item {
	var obj Item

	for _, e := range r.Events {

		obj.Meets = append(obj.Meets, e.Summary+" ("+e.HangoutLink+")")
	}

	for _, i := range r.Issues {
		var task string
		task += i.Key + " " + i.Fields.Summary + " (https://cureety.atlassian.net/browse/" + i.Key + ")"
		if strings.ToLower(i.Fields.Status.Name) == "in progress" {
			obj.Tasks.Progress = append(obj.Tasks.Progress, task)
		}

		if strings.ToLower(i.Fields.Status.Name) == "done" {
			obj.Tasks.Done = append(obj.Tasks.Done, task)
		}
	}

	for _, c := range r.Contributions {
		obj.Contributions = append(obj.Contributions, Contribution{
			Action: c.ActionName,
			Name:   strContr(c),
		})
	}

	obj.Hours = r.Hours
	obj.Note = r.Note
	obj.Date = r.Date.Format("2006-01-02")

	return obj
}
