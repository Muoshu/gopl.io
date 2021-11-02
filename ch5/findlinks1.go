package ch5

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
)

func VisitMain() {

	url := "http://gopl.io"
	doc, err := html.Parse(fetch(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks:%v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}

func OutlineMain() {
	url := "http://gopl.io"
	doc, err := html.Parse(fetch(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline:#{err}\n")
		os.Exit(1)
	}
	outline(nil, doc)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links

}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func fetch(url string) *bytes.Reader {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	return bytes.NewReader(b)
}
