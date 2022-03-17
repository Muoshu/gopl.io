package main

import (
	"awesomeProject/ch5"
	"fmt"
	"log"
)

func main() {
	worklist := make(chan []string)
	go func() {
		worklist <- []string{"http://gopl.io"}
	}()
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}

}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := ch5.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
