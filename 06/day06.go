package main

import (
	"fmt"
	"log"
	"os"
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

	cpy1 := make([]string, len(input))
	copy(cpy1, input)
	findTiles(cpy1)

	cpy2 := make([]string, len(input))
	copy(cpy2, input)
	countObstructions(cpy2)
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
func findStart(input []string) (int, int) {
	for x := range input {
		line := input[x]
		for y := range len(line) {
			if line[y] == '^' {
				//fmt.Printf("%v %v\n", x, y)
				return x, y
			}
		}
	}
	fmt.Printf("%v %v\n", 0, 0)
	return 0, 0
}

type dirs struct {
	dx, dy int
}

func isObstructed(input []string) bool {
	directions := []dirs{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	dir := 0
	x, y := findStart(input)
	for true {
		nx, ny := x+directions[dir].dx, y+directions[dir].dy
		if nx < 0 || nx >= len(input) || ny < 0 || ny >= len(input[nx]) {
			return false
		}
		char := input[nx][ny]
		//fmt.Printf("char: %v (%v,%v)", string(char), nx, ny)
		switch char {
		case '#':
			dir = (dir + 1) % 4
		case '.':
			out := []rune(input[nx])
			out[ny] = 1 << (dir)
			input[nx] = string(out)
			x, y = nx, ny
		case '^':
			out := []rune(input[nx])
			out[ny] = 1 << (dir)
			input[nx] = string(out)
			x, y = nx, ny
		default:
			//fmt.Printf("(%d & %d -> %d) == %d -> %v\n", int(char), 1<<dir, int(char)&(1<<(dir)), 1<<(dir), (int(char)&(1<<(dir))) == 1<<(dir))

			if (int(char) & (1 << (dir))) == 1<<(dir) {
				return true
			}
			newchar := rune((char | (1 << (dir))))
			//fmt.Printf("%v -> %v\n", char, newchar)
			out := []rune(input[nx])
			out[ny] = newchar
			input[nx] = string(out)
			x, y = nx, ny
		}
	}
	return false
}

func countObstructions(input []string) {
	defer timerHighRes("Day 6 Part 2 countObstructions")()
	count := 0
	cpy1 := make([]string, len(input))
	copy(cpy1, input)
	if isObstructed(cpy1) {
		count += 1
	}
	for x := range input {
		line := input[x]
		for y := range len(line) {
			if line[y] == '.' {
				cpy := make([]string, len(input))
				copy(cpy, input)
				out := []rune(cpy[x])
				out[y] = '#'
				cpy[x] = string(out)
				if isObstructed(cpy) {
					count += 1
				}
			}
		}
	}
	fmt.Printf("Count: %v\n", count)

}

func findTiles(input []string) {
	defer timerHighRes("Day 6 Part 1 fildTiles")()
	x, y := findStart(input)
	directions := []dirs{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	dir := 0
	//fmt.Printf("Dimensions: %v, %v\n", len(input), len(input[0]))
	count := 1
	for true {
		nx, ny := x+directions[dir].dx, y+directions[dir].dy
		if nx < 0 || nx >= len(input) || ny < 0 || ny >= len(input[nx]) {
			fmt.Printf("Count: %v\n", count)
			return
		}
		char := input[nx][ny]
		//fmt.Printf("char: %v (%v,%v)", string(char), nx, ny)
		switch char {
		case '#':
			dir = (dir + 1) % 4
		case '.':
			out := []rune(input[nx])
			out[ny] = 'X'
			input[nx] = string(out)
			x, y = nx, ny
			count++
		case 'X':
			x, y = nx, ny
		case '^':
			x, y = nx, ny
		}
		//fmt.Printf("\n")
	}
	fmt.Printf("Count: %v\n", count)
}
