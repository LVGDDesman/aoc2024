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
	inputraw, err := os.ReadFile("task.txt")
	//inputraw, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(inputraw)
	numbers := []int{}
	for _, snum := range strings.Split(input, " ") {
		num, err := strconv.Atoi(snum)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	nmap := make(map[int]int)
	for _, snum := range strings.Split(input, " ") {
		num, err := strconv.Atoi(snum)
		if err != nil {
			log.Fatal(err)
		}
		nmap[num] += 1
	}

	evolveStonesSlow(numbers, 25)
	//evolveStonesSlow(numbers, 75) // don't

	evolveMoreStones(nmap, 25, "1")
	evolveMoreStones(nmap, 75, "2")

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

func evolveMoreStones(input map[int]int, evolutions int, task string) {
	defer timerHighRes("Day 11 Part " + task + " evolveStones")()
	count := 0
	for _, v := range input {
		count += v
	}
	//fmt.Printf("%v\n", input)
	for range evolutions {
		newInput := make(map[int]int)
		for i, v := range input {
			if i == 0 {
				newInput[1] += v
			} else {
				t := i
				vlen := 0
				for t != 0 {
					t /= 10
					vlen += 1
				}
				t = i
				if vlen%2 == 0 {
					nval := 0
					power := 1
					for range vlen / 2 {
						nval = nval + t%10*power
						power *= 10
						t /= 10
					}

					newInput[t] += v
					newInput[nval] += v
					count += v

				} else {
					newInput[i*2024] += v
				}
			}
		}
		input = newInput
		//fmt.Printf("%v\n", input)
	}
	fmt.Printf("Count: %v\n", count)
}

func evolveStonesSlow(input []int, evolutions int) {
	defer timerHighRes("Day 11 Part 1 evolveStones")()

	for i := range evolutions {
		for i, v := range input {
			if v == 0 {
				input[i] = 1
			} else {
				t := v
				vlen := 0
				for t != 0 {
					t /= 10
					vlen += 1
				}
				if vlen%2 == 0 {
					nval := 0
					power := 1
					for range vlen / 2 {
						nval = nval + v%10*power
						power *= 10
						v /= 10
					}
					input[i] = v
					input = append(input, nval)

				} else {
					input[i] *= 2024
				}
			}
		}
		fmt.Printf("%v\n", i)
		//fmt.Printf("%v\n", input)
	}
	fmt.Printf("Count: %v\n", len(input))
}
