package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	f, err := os.ReadFile("task.txt")
	//f, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", findmul(string(f)))
	fmt.Printf("%v\n", finddomul(string(f)))
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

func finddomul(input string) int {
	defer timerHighRes("Day 3 2")()
	// graveayrd
	// (don't\(\)(?:(do\(\))|(don't\(\)))*do\(\))?(?:(do\(\))|(don't\(\)))*
	// (?:don't\(\))?
	// (^|do\(\)).*?((mul\((\d+),(\d+)\)).*?)*(?:don't\(\)|$)
	// (?:^|do\(\))(?:.*?(?:don't\(\).*?do\(\))?)(?:mul\((\d+),(\d+)\))
	re := regexp.MustCompile(`(do\(\))|(don't\(\))|(mul\((\d+),(\d+)\))`)
	count := 0
	do := true
	for _, submatches := range re.FindAllStringSubmatch(input, -1) {
		if submatches[0] == "don't()" {
			do = false
		} else if submatches[0] == "do()" {
			do = true
		} else if do {
			a, err := strconv.Atoi(submatches[4])
			b, err := strconv.Atoi(submatches[5])

			if err != nil {
				panic(err)
			}
			count += a * b
		}

	}
	return count
}

func findmul(input string) int {
	defer timerHighRes("Day 3 1")()
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	count := 0
	for _, submatches := range re.FindAllStringSubmatch(input, -1) {
		a, err := strconv.Atoi(submatches[1])
		b, err := strconv.Atoi(submatches[2])
		if err != nil {
			panic(err)
		}
		count += a * b
	}
	return count
}
