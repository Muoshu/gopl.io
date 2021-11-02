package ch4

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Iterms     []*Issue
}

type Issue struct {
	Number  int
	HTMLURL string `json:"html_url"`
	Title   string
	User    *User
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssue(term []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(term, " "))
	s := IssueURL + "?q=" + q
	fmt.Println(s)
	resp, err := http.Get(s)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func SearchIssueMain() {
	terms := []string{"repo:golang/go is:open json decoder"}
	result, err := SearchIssue(terms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println(result.Iterms)
	for _, item := range result.Iterms {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
