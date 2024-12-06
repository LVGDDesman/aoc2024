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
	testPages02(input)
	rectify02(input)
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

func rectify02(input []string) {
	defer timerHighRes("Day 5 2 rectify optimized")()
	count := 0
	isPart1 := true
	rules := make(map[int]map[int]bool)
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
			if rules[a] == nil {
				rules[a] = make(map[int]bool)
			}
			rules[a][b] = true
		} else {
			pages := strings.Split(line, ",")
			if pos := findError02(rules, pages); pos != -1 {
				//fmt.Printf("rectify %v at %v\n", line, pos)
				for true {
					pages[pos], pages[pos-1] = pages[pos-1], pages[pos]
					//fmt.Printf("try: %v at %v\n", pages, pos)

					if pos = findError02(rules, pages); pos == -1 {
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

func rectify(input []string) {
	defer timerHighRes("Day 5 2 rectify not optimized")()
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
			if pos := findError(rules, pages); pos != -1 {
				//fmt.Printf("rectify %v at %v\n", line, pos)
				for true {
					pages[pos], pages[pos-1] = pages[pos-1], pages[pos]
					//fmt.Printf("try: %v at %v\n", pages, pos)

					if pos = findError(rules, pages); pos == -1 {
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

func testPages02(input []string) {
	defer timerHighRes("Day 5 1 testPages optimized")()
	count := 0
	isPart1 := true
	rules := make(map[int]map[int]bool)

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
			if rules[a] == nil {
				rules[a] = make(map[int]bool)
			}
			rules[a][b] = true
		} else {
			pages := strings.Split(line, ",")
			middle, err := strconv.Atoi(pages[len(pages)/2])
			if err != nil {
				panic(err)
			}
			if findError02(rules, pages) == -1 {
				count += middle
			}

		}

	}
	fmt.Printf("%v\n", count)
}

func testPages(input []string) {
	defer timerHighRes("Day 5 1 testPages not optimized")()
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
			if findError(rules, pages) == -1 {
				count += middle
			}

		}

	}
	fmt.Printf("%v\n", count)
}

func findError02(rules map[int]map[int]bool, pages []string) int {
	oldpages := make([]int, 0, len(pages))
	for x, pagestr := range pages {
		page, err := strconv.Atoi(pagestr)
		if err != nil {
			panic(err)
		}
		for _, oldpage := range oldpages {

			if rules[page] != nil && rules[page][oldpage] {
				return x
			}
		}
		oldpages = append(oldpages, page)
	}
	return -1
}

func findErrorO1(rules map[int]map[int]bool, pages []string) int {
	//oldpages := []int{}
	checked := make(map[int]bool)
	for x, pagestr := range pages {
		page, err := strconv.Atoi(pagestr)
		if err != nil {
			panic(err)
		}
		if dep, exists := rules[page]; exists {
			for d := range dep {
				if checked[d] {
					return x
				}
			}
		}
		checked[page] = true
	}
	return -1
}

func findErrorO0(rules [][]int, pages []string) int {
	checked := make(map[int]bool)
	for x, pagestr := range pages {
		page, err := strconv.Atoi(pagestr)
		if err != nil {
			panic(err)
		}
		for _, rule := range rules {
			if rule[0] == page && checked[rule[1]] {
				return x
			}
		}
		checked[page] = true
	}
	return -1
}

func findError(rules [][]int, pages []string) int {
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
