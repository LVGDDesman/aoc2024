package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	inputraw, err := os.ReadFile("task.txt")
	//inputraw, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := [][]byte{}
	length := 0
	input = append(input, []byte{})
	for _, char := range inputraw {
		if char == '\n' {
			length++
			input = append(input, []byte{})
		} else {
			input[length] = append(input[length], byte(char))
		}
	}

	findTrails(input)
	findAllTrails(input)
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

type dirs struct {
	dx, dy int
}

func findAllTrails(input [][]byte) {
	defer timerHighRes("Day 10 Part 2 findAllTrails")()
	count := 0
	for x := range input {
		for y := range input[x] {
			if input[x][y] == '0' {
				findAllTrailHead(input, x, y, '1', &count)
				//fmt.Printf("found Start: %v, %v now %v\n", x, y, count)
			}
		}
	}
	fmt.Printf("Count: %v\n", count)
}

func findAllTrailHead(input [][]byte, x int, y int, currIndex byte, count *int) {
	directions := []dirs{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	//fmt.Printf("Searching for %v around %v, %v\n", currIndex, x, y)
	for _, dir := range directions {
		nx := x + dir.dx
		ny := y + dir.dy
		if nx < 0 || nx >= len(input) || ny < 0 || ny >= len(input[nx]) {
			//fmt.Printf("  ->(%v, %v) == %v Outof Bounds\n", nx, ny, currIndex)
			continue
		}
		if currIndex == '9' && input[nx][ny] == '9' {
			//fmt.Printf("  ->(%v, %v) == %v FINAL HIT\n", nx, ny, currIndex)
			*count += 1
			continue
		}
		if input[nx][ny] == currIndex {
			//fmt.Printf("  ->(%v, %v) == %v HIT\n", nx, ny, currIndex)
			findAllTrailHead(input, nx, ny, currIndex+1, count)
		} else {
			//fmt.Printf("  ->WRONG (%v %v) != %v, %v\n", nx, ny, currIndex, input[nx][ny])

		}
	}
}

func findTrails(input [][]byte) {
	defer timerHighRes("Day 10 Part 1 findTrails")()
	count := 0
	for x := range input {
		for y := range input[x] {
			if input[x][y] == '0' {
				findTrailHead(input, x, y, '1', &count, make(map[int]map[int]bool))
				//fmt.Printf("found Start: %v, %v now %v\n", x, y, count)
			}
		}
	}
	fmt.Printf("Count: %v\n", count)
}

func findTrailHead(input [][]byte, x int, y int, currIndex byte, count *int, heads map[int]map[int]bool) {
	directions := []dirs{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	//fmt.Printf("Searching for %v around %v, %v\n", currIndex, x, y)
	for _, dir := range directions {
		nx := x + dir.dx
		ny := y + dir.dy
		if nx < 0 || nx >= len(input) || ny < 0 || ny >= len(input[nx]) {
			//fmt.Printf("  ->(%v, %v) == %v Outof Bounds\n", nx, ny, currIndex)
			continue
		}
		if currIndex == '9' && input[nx][ny] == '9' {
			if heads[nx] == nil {
				heads[nx] = make(map[int]bool)
			}
			if !heads[nx][ny] {
				heads[nx][ny] = true
				//fmt.Printf("  ->(%v, %v) == %v FINAL HIT\n", nx, ny, currIndex)
				*count += 1
			}
			continue
		}
		if input[nx][ny] == currIndex {
			//fmt.Printf("  ->(%v, %v) == %v HIT\n", nx, ny, currIndex)
			findTrailHead(input, nx, ny, currIndex+1, count, heads)
		} else {
			//fmt.Printf("  ->WRONG (%v %v) != %v, %v\n", nx, ny, currIndex, input[nx][ny])

		}
	}
}
