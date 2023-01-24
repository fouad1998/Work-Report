package calendar

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os/exec"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	var authCode string
	var err error = exec.Command("xdg-open", authURL).Start()
	if err != nil {
		fmt.Printf("Please go to the following link in your browser:  \n%s\n", authURL)
		fmt.Println("The copy authorization code from url response and paste in here")
		if _, err := fmt.Scan(&authCode); err != nil {
			return nil, err
		}
	} else {
		listener, err := net.Listen("tcp", ":3232")
		if err != nil {
			return nil, err
		}

		wg := sync.WaitGroup{}

		wg.Add(1)
		go func() {
			done := false
			server := http.Server{}
			http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				authCode = r.URL.Query().Get("code")
				io.WriteString(w, `
				<!DOCTYPE html>
				<html lang="en">

				<head>
				   <meta charset="UTF-8">
				   <meta http-equiv="X-UA-Compatible" content="IE=edge">
				   <meta name="viewport" content="width=device-width, initial-scale=1.0">
				</head>

				<body>
				  <h1 style="text-align: center; font-family: sans-serif; color: #364149;padding: 150px 0px;">
				      Thank you. Return to app
				   </h1>
				</body>

				</html>
			`)

				if !done {
					done = true
					wg.Done()
					go func() {
						time.Sleep(5 * time.Second)
						server.Close()
					}()
				}
			}))

			server.Serve(listener)
		}()

		wg.Wait()
	}

	return config.Exchange(context.TODO(), authCode)
}
