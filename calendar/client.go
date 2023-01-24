package calendar

import (
	"context"

	"github.com/fouad1998/work-reporter/env"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func (c *Calendar) getClient() error {
	ctx := context.Background()
	if err := c.config.Read(); err != nil {
		return err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON([]byte(env.Conf.GoogleApp), calendar.CalendarReadonlyScope)
	if err != nil {
		return err
	}

	config.RedirectURL = "http://localhost:3232"
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.

	if c.config.GToken == nil {
		token, err := getTokenFromWeb(config)
		if err != nil {
			return err
		}

		c.config.GToken = token
		c.config.Save()
	}

	client := config.Client(context.Background(), c.config.GToken)

	service, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return err
	}

	c.service = service

	return nil
}
