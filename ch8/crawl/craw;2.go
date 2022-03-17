package main

import (
	"awesomeProject/ch5"
	"fmt"
	"log"
)

func main() {
	workList := make(chan []string)
	unSeenLinks := make(chan string)

	go func() { workList <- []string{"http://gopl.io"} }()
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unSeenLinks {
				foundLinks := Crawl(link)
				go func() { workList <- foundLinks }()
			}
		}()
	}
	seen := make(map[string]bool)
	for list := range workList {
		fmt.Println(len(list))
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unSeenLinks <- link
			}
		}
	}

}

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := ch5.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
