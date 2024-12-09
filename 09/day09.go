package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	input, err := os.ReadFile("task.txt")
	//input, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fragmentAndCheckSum(input)
	defragmentAndCheckSum(input)
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

func defragmentAndCheckSum(input []byte) {
	defer timerHighRes("Day 9 Part 2 defragmentAndCheckSum")()
	// var result strings.Builder
	counts := make([]int, len(input))
	for i, v := range input {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal(err)
		}
		counts[i] = n
	}

	res := 0
	pos := 0
	visited := make([]bool, len(counts))
	for i, v := range counts {
		last := len(counts)
		if last%2 == 1 {
			last--
		}
		if i%2 == 0 {
			if visited[i] {
				pos += v
				// for ; v != 0; v-- {
				// 	result.Write([]byte("."))
				// }
				continue
			}
			for ; v != 0; v-- {
				// result.Write([]byte(strconv.Itoa(i / 2)))
				res += i / 2 * pos
				pos++
			}
		} else {
			for ; last >= i; last -= 2 {
				r := counts[last]
				if visited[last] || r == 0 || r > v {
					continue
				}
				visited[last] = true
				v -= r
				for ; r != 0; r-- {
					// result.Write([]byte(strconv.Itoa(last / 2)))
					res += last / 2 * pos
					pos++
				}
			}
			pos += v
			// for ; v != 0; v-- {
			// 	result.Write([]byte("."))
			// }

		}
	}
	//fmt.Printf("%v (as value: %v)\n", result.String(), res)
	fmt.Printf("Checksum: %v\n", res)
}

func fragmentAndCheckSum(input []byte) {
	defer timerHighRes("Day 9 Part 1 fragmentAndCheckSum")()
	//var result strings.Builder
	counts := make([]int, len(input))
	for i, v := range input {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal(err)
		}
		counts[i] = n
	}
	last := len(counts)
	if last%2 == 1 {
		last--
	}
	res := 0
	pos := 0
	for x, i := range counts {
		if x%2 == 0 {
			for ; i != 0; i-- {
				//result.Write([]byte(strconv.Itoa(x / 2)))
				res += x / 2 * pos
				pos++
			}
		} else {
			if x >= last {
				break
			}
			for i != 0 {
				r := counts[last]
				if r <= i {
					i -= r
					for ; r != 0; r-- {
						//result.Write([]byte(strconv.Itoa(last / 2)))
						res += last / 2 * pos
						pos++
					}
					counts[last] = 0
					last -= 2
				} else {
					r -= i
					counts[last] = r
					for ; i != 0; i-- {
						//result.Write([]byte(strconv.Itoa(last / 2)))
						res += last / 2 * pos
						pos++
					}
					i = 0
				}
			}
		}
	}
	//fmt.Printf("%v (as value: %v)\n", result.String(), res)
	fmt.Printf("Checksum: %v\n", res)
}
