package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("task.txt")
	//f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//defer timerHighRes("Day 1 1")()
	defer timerHighRes("Day 1 2")()

	scanner := bufio.NewScanner(f)
	var left = []int{}
	var right = []int{}
	var i = 1
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "   ")
		l, err := strconv.Atoi(split[0])
		r, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		left = append(left, l)
		right = append(right, r)
		i = i + 1
	}
	sort.Slice(left, func(i, j int) bool {
		return left[i] > left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] > right[j]
	})
	result1 := countDifferences(left, right)
	fmt.Printf("Difference: %v\n", result1)
	//result2 := similarity(left, right)
	//fmt.Printf("Similarity: %v\n", result2)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
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

func countDifferences(left, right []int) int {
	var result int
	for i := 0; i < len(left); i++ {
		res := left[i] - right[i]
		//fmt.Printf("%i - %i = %i\n", left[i], right[i], res)
		if res < 0 {
			result = result - res
		} else {
			result = result + res
		}
	}
	return result
}

func similarity(left, right []int) int {
	var result int
	for i := 0; i < len(left); i++ {
		var count int
		value := left[i]
		for r := 0; r < len(right); r++ {
			if value == right[r] {
				count = count + 1
			} else if value > right[r] {
				break
			}
		}
		result += value * count

	}
	return result
}
