package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"learning/go-programming-language/chapter4/4.5/github"
)

const temp1 = `{{.TotalCount}} issues:
{{range .Items}}----------------------------
Number:{{.Number}}
User:{{.User.Login}}
Title:{{.Title | printf "%.64s"}}
Age:{{.CreateAt | daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.
	New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(temp1))

func Report() {
	result, err := github.SearchIssues([]string{"repo:golang/go", "is:open json decoder"})
	if err != nil {
		log.Fatal(err)
	}
	if err := issuesList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
