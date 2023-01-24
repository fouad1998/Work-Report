package jira

func (j *Jira) SetToken(token, email string) error {
	j.config.JiraToken = token
	j.config.JiraEmail = email

	return j.config.Save()
}
