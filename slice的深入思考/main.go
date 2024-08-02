package main

import (
	"fmt"
	"log"
)

func main() {
	ch1 := make(chan int, 10)
	ch1 <- 1
	agree := <-ch1

	// agree := <-ch1
	fmt.Println("接受", agree)

	Test_slice()
}

var s2 = []int{}
var map1 = map[string]string{}

func Test_slice() {
	// s := make([]int,10)
	// s = append(s, 10)
	// log.Fatalf("s: %v, len of s: %d, cap of s: %d,s,len(s),cap(s)", s, len(s), cap(s)) //len 11 cap: 20
	// s := make([]int,0,10)
	// s = append(s, 10)
	// log.Fatalf("s: %v, len of s: %d, cap of s: %d,", s, len(s), cap(s)) //len 1 cap: 10
	// s := make([]int, 10, 11)
	// s = append(s, 10)
	// log.Fatalf("s: %v, len of s: %d, cap of s: %d,", s, len(s), cap(s)) //len 11 cap: 11
	// s := make([]int, 10, 12)
	// s1 := s[8:]
	// log.Fatalf("s: %v, len of s: %d, cap of s: %d,", s1, len(s1), cap(s1)) // [0 0], len of s: 2, cap of s: 4,
	s := make([]int, 10, 12)
	s1 := s[8:9]
	log.Fatalf("s: %v, len of s: %d, cap of s: %d,", s1, len(s1), cap(s1)) //[0], len of s: 1, cap of s: 4,
}
