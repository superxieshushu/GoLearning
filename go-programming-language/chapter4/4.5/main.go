package main

import "learning/go-programming-language/chapter4/4.5/github"

func main() {
	//struct2Json()
	//json2Struct()
	github.ListIssues([]string{"repo:golang/go", "is:open json decoder"})
}
