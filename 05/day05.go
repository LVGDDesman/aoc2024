package main

import (
	//"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.ReadFile("task.txt")
	//f, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(f), "\n")

	testPages(input)
	rectify(input)
}

func timerHighRes(name string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		// Print duration with nanosecond precision
		fmt.Printf("%s took %d ns (%.6f ms)\n",
			name,
			elapsed.Nanoseconds(),
			float64(elapsed.Nanoseconds())/1_000_000.0)
	}
}

func rectify(input []string) {
	defer timerHighRes("Day 5 2 rectify")()
	count := 0
	isPart1 := true
	rules := [][]int{}
	for i := range input {
		line := input[i]
		if line == "" {
			isPart1 = false
		} else if isPart1 == true {
			a, err := strconv.Atoi(strings.Split(line, "|")[0])
			b, err := strconv.Atoi(strings.Split(line, "|")[1])
			if err != nil {
				panic(err)
			}
			rules = append(rules, []int{a, b})

		} else {
			pages := strings.Split(line, ",")
			if pos := condMiddle(rules, pages); pos != -1 {
				//fmt.Printf("rectify %v at %v\n", line, pos)
				for true {
					pages[pos], pages[pos-1] = pages[pos-1], pages[pos]
					//fmt.Printf("try: %v at %v\n", pages, pos)

					if pos = condMiddle(rules, pages); pos == -1 {
						//fmt.Printf("rectified by: %v)\n", pages)
						middle, err := strconv.Atoi(pages[len(pages)/2])
						if err != nil {
							panic(err)
						}
						count += middle
						break
					}
				}
			}
		}
	}
	fmt.Printf("%v\n", count)
}

func testPages(input []string) {
	defer timerHighRes("Day 5 1 testPages")()
	count := 0
	isPart1 := true
	rules := [][]int{}
	for i := range input {
		line := input[i]
		if line == "" {
			isPart1 = false
		} else if isPart1 == true {
			a, err := strconv.Atoi(strings.Split(line, "|")[0])
			b, err := strconv.Atoi(strings.Split(line, "|")[1])
			if err != nil {
				panic(err)
			}
			rules = append(rules, []int{a, b})

		} else {
			pages := strings.Split(line, ",")
			middle, err := strconv.Atoi(pages[len(pages)/2])
			if err != nil {
				panic(err)
			}
			if condMiddle(rules, pages) == -1 {
				count += middle
			}

		}

	}
	fmt.Printf("%v\n", count)
}

func condMiddle(rules [][]int, pages []string) int {
	oldpages := []int{}
	for x := range pages {
		page, err := strconv.Atoi(pages[x])
		if err != nil {
			panic(err)
		}
		for y := range oldpages {
			oldpage := oldpages[y]
			for z := range rules {
				rule := rules[z]
				if rule[0] == page && rule[1] == oldpage {
					return x
				}
			}
		}
		oldpages = append(oldpages, page)
	}
	return -1
}
