package main

import (
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


	findCalibrationResult(input)
	findCalibrationResult2(input)
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
func findCalibrationResult2(input []string) {
	defer timerHighRes("Day 7 2 findCalibrationResult2")()
	var count int64 = 0
	for _, line := range input {
		parts := strings.Split(line, ": ")
		result, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbersstr := strings.Split(parts[1], " ")
		var numbers []int64
		for _,number := range numbersstr {
			n, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, n)
		}

		if isSolving2(result, numbers, 1, numbers[0]) {
			count += result
		}
	}
	fmt.Printf("Count: %v\n", count)
}

func isSolving2(result int64, numbers []int64, index int, curres int64) bool {
	if index >= len(numbers) {
		if curres == result {
			return true
		} else {
			return false
		}
	}
	if curres > result {
		return false
	}
	concat := curres
	i := numbers[index]
	for  i != 0 {
        i /= 10
        concat *=10
    }
	concat += numbers[index]
	return isSolving2(result, numbers, index+1, curres + numbers[index]) ||
		   isSolving2(result, numbers, index+1, curres * numbers[index]) ||
		   isSolving2(result, numbers, index+1, concat)
}

func findCalibrationResult(input []string) {
	defer timerHighRes("Day 7 1 findCalibrationResult")()
	count := 0
	for _, line := range input {
		parts := strings.Split(line, ": ")
		result, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		numbersstr := strings.Split(parts[1], " ")
		var numbers []int
		for _,number := range numbersstr {
			n, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, n)
		}

		if isSolving(result, numbers, 1, numbers[0]) {
			count += result
		}
	}
	fmt.Printf("Count: %v\n", count)
}

func isSolving(result int, numbers []int, index int, curres int) bool {
	if index >= len(numbers) {
		if curres == result {
			return true
		} else {
			return false
		}
	}
	if curres > result {
		return false
	}
	return isSolving(result, numbers, index+1, curres + numbers[index]) ||
		   isSolving(result, numbers, index+1, curres * numbers[index])
}