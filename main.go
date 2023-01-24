package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fouad1998/work-reporter/calendar"
	"github.com/fouad1998/work-reporter/gitlab"
	"github.com/fouad1998/work-reporter/jira"
	"github.com/fouad1998/work-reporter/reporter"
)

var title string = `

      __     __     ______     ______     __  __        ______     ______     ______   ______     ______     ______  
     /\ \  _ \ \   /\  __ \   /\  == \   /\ \/ /       /\  == \   /\  ___\   /\  == \ /\  __ \   /\  == \   /\__  _\ 
     \ \ \/ ".\ \  \ \ \/\ \  \ \  __<   \ \  _"-.     \ \  __<   \ \  __\   \ \  _-/ \ \ \/\ \  \ \  __<   \/_/\ \/ 
      \ \__/".~\_\  \ \_____\  \ \_\ \_\  \ \_\ \_\     \ \_\ \_\  \ \_____\  \ \_\    \ \_____\  \ \_\ \_\    \ \_\ 
       \/_/   \/_/   \/_____/   \/_/ /_/   \/_/\/_/      \/_/ /_/   \/_____/   \/_/     \/_____/   \/_/ /_/     \/_/ 
                                                                                                                     

`

var options = []string{
	"Configure your jira token",
	"Configure your gitlab token",
	"Generate your daily report",
	"Generate your missed reports",
	"Export as excel",
}

func main() {
	report := reporter.Reporter{}
	calendar := calendar.Calendar{}
	jira := jira.Jira{}
	gitlab := gitlab.Gitlab{}

	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Printf("%s", title)

		fmt.Println(strings.Repeat("\n", 4) + "You have the following options \n")

		for index, option := range options {
			fmt.Println(index+1, "->", option)
		}

		fmt.Print(strings.Repeat("\n", 2) + "Choose an option: ")

		var option int
		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			fallthrough

		case 2:
			var token string
			fmt.Print(strings.Repeat("\n", 4))
			if option == 1 {
				fmt.Println("Please visite the following link in order to get your Jira token. (https://id.atlassian.com/manage-profile/security/api-tokens)")
			} else {
				fmt.Println("Please visite the following link in order to get your Gitlab token (https://gitlab.com/-/profile/personal_access_tokens)")
			}

			fmt.Print("Token: ")
			if _, err := fmt.Scanf("%s", &token); err != nil {
				log.Fatalln(err)
			}

			if option == 1 {
				fmt.Print("Jira email: ")
				var email string
				if _, err := fmt.Scanf("%s", &email); err != nil {
					log.Fatalln(err)
				}

				jira.SetToken(token, email)
				continue
			}

			gitlab.SetToken(token)

		case 3:
			var hours int
			var note string
			fmt.Print(strings.Repeat("\n", 4))
			has, err := report.Has(time.Now())
			if err != nil {
				log.Fatalln(err)
			}

			if has {
				var response string
				fmt.Println("You have already reported for today, if you proceed you will lost the previous report.")
				fmt.Print("Are you sure to continue (y/n)? ")
				fmt.Scanf("%s", &response)
				response = strings.ToLower(response)
				if !(response == "y" || response == "yes") {
					break
				}
			}

			fmt.Print("Number of hours have you worked today?: ")
			if _, err := fmt.Scanf("%d", &hours); err != nil {
				log.Fatalln(err)
			}

			fmt.Println("Add some other work: ")
			for {
				var line string
				fmt.Scanf("%s", &line)
				if line == "" {
					break
				}

				note += line + "\n"
			}

			fmt.Println("Loading...")
			events, err := calendar.GetTodayEvents()
			if err != nil {
				log.Fatalln(err)
			}

			issues, err := jira.GetTodayTasks()
			if err != nil {
				log.Fatalln(err)
			}

			contributions, err := gitlab.GetTodayContribution()
			if err != nil {
				log.Fatalln(err)
			}

			report.Add(&reporter.Report{
				Date:          time.Now(),
				Issues:        issues,
				Events:        events,
				Contributions: contributions,
				Hours:         hours,
				Note:          note,
			})

		case 4:
			fallthrough

		case 5:
			var input string
			fmt.Print(strings.Repeat("\n", 4) + "Please type the date month following this format (YYYY-MM) ?: ")
			if _, err := fmt.Scanf("%s", &input); err != nil {
				log.Fatalln(err)
			}

			date, err := time.Parse("2006-01", input)
			if err != nil {
				log.Fatalln(err)
			}

			if option == 5 {
				if err := report.Excel(date); err != nil {
					log.Fatalln(err)
				}

				continue
			}

			if time.Now().Before(date) {
				fmt.Println("You can not get the report for future dates")
				continue
			}

			for i := 0; i < 31; i++ {
				workingDate := date.Add(time.Duration(i*24) * (time.Hour))
				if workingDate.After(time.Now()) {
					continue
				}

				if workingDate.Weekday() == time.Saturday || workingDate.Weekday() == time.Sunday {
					continue
				}

				has, err := report.Has(workingDate)
				if err != nil {
					log.Fatalln(err)
				}

				if has {
					continue
				}

				fmt.Println("Loading report for " + workingDate.Format("2006-01-02"))
				events, err := calendar.GetEvents(workingDate)
				if err != nil {
					log.Fatalln(err)
				}

				issues, err := jira.GetTasks(workingDate)
				if err != nil {
					log.Fatalln(err)
				}

				contributions, err := gitlab.GetContributions(workingDate)
				if err != nil {
					log.Fatalln(err)
				}

				report.Add(&reporter.Report{
					Date:          workingDate,
					Issues:        issues,
					Events:        events,
					Contributions: contributions,
					Hours:         8,
				})
			}
		}
	}
}
