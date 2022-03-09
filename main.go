package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	a := []int{6, 2, 10, 3, 4, 5, 1}
	result := merge(a, 0, 6)
	fmt.Println(result)
	fmt.Println(a)
}

func merge(a []int, s, e int) int {
	if s < e {
		mid := (s + e) / 2
		l := merge(a, s, mid)
		r := merge(a, mid+1, e)
		m := mergeCount(a, s, mid, e)
		return l + r + m
	}
	return 0
}

func mergeCount(a []int, s, m, e int) int {
	size := len(a)
	temp := make([]int, size)
	i, j, k, count := s, m+1, s, 0
	for i := s; i <= e; i++ {
		temp[i] = a[i]
	}
	for i <= m && j <= e {
		if temp[i] < temp[j] {
			a[k] = temp[i]
			k++
			i++
		} else {
			a[k] = temp[j]
			k++
			j++
			count = m - i + 1
		}

	}
	for i <= m {
		a[k] = temp[i]
		k++
		i++
	}
	for j <= e {
		a[k] = temp[j]
		k++
		j++
	}
	return count
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
