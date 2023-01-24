package reporter

type Contribution struct {
	Action string `json:"action"`
	Name   string `json:"name"`
}

type Item struct {
	Date          string         `json:"date"`
	Meets         []string       `json:"meets"`
	Contributions []Contribution `json:"contributions"`
	Tasks         struct {
		Progress []string
		Done     []string
	} `json:"tasks"`
	Note  string `json:"note"`
	Hours int    `json:"hours"`
}

type Reporter struct {
	Items []Item
}
