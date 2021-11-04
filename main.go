package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func test() (i int) {
	defer func() { i++ }()
	return 1
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()

	time.Sleep(5 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s),", msg, time.Since(start))
	}

}
