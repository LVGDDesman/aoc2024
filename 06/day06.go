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

	cpy1 := deepcopy(input)
	findTiles(cpy1)

	cpy2 := deepcopy(input)
	countObstructions(cpy2)
}

func printField(input [][]byte) {
	for x := range input {
		for y := range input[x] {
			fmt.Printf("%v", string(input[x][y]))
		}
		fmt.Printf("\n")
	}
}
func deepcopy(input [][]byte) [][]byte {
	cpy1 := make([][]byte, len(input))
	for i, row := range input {
		cpy1[i] = make([]byte, len(row))
		copy(cpy1[i], row)
	}
	return cpy1
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
func findStart(input [][]byte) (int, int) {
	for x := range input {
		line := input[x]
		for y := range len(line) {
			if line[y] == '^' {
				//fmt.Printf("Start: %v %v\n", x, y)
				return x, y
			}
		}
	}
	fmt.Printf("HALLO %v %v\n", 0, 0)
	return 0, 0
}

type dirs struct {
	dx, dy int
}

func isObstructed(input [][]byte, x int, y int, directions []dirs, count *int) {
	dir := 0
	//fmt.Printf("%v, %v\n", x, y)
	for true {
		nx, ny := x+directions[dir].dx, y+directions[dir].dy
		if nx < 0 || nx >= len(input) || ny < 0 || ny >= len(input[nx]) {
			return
		}
		char := input[nx][ny]
		//fmt.Printf("char: %v (%v,%v)", string(char), nx, ny)
		switch char {
		case '#':
			dir = (dir + 1) % 4
		case '.':
			x, y = nx, ny
			input[nx][ny] = 1 << dir
		case '^':
			x, y = nx, ny
			input[nx][ny] = 1 << dir
		default:
			//fmt.Printf("(%d & %d -> %d) == %d -> %v\n", int(char), 1<<dir, int(char)&(1<<(dir)), 1<<(dir), (int(char)&(1<<(dir))) == 1<<(dir))

			if (int(char) & (1 << (dir))) == 1<<(dir) {
				*count++
				return
			}
			newchar := char | (1 << dir)
			//fmt.Printf("%v -> %v\n", char, newchar)
			input[nx][ny] = newchar
			x, y = nx, ny
		}
	}
	return
}

func countObstructions(input [][]byte) {
	defer timerHighRes("Day 6 Part 2 countObstructions")()

	directions := []dirs{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	sx, sy := findStart(input)
	count := 0
	go isObstructed(deepcopy(input), sx, sy, directions, &count)
	for x := range input {
		line := input[x]
		for y := range len(line) {
			if line[y] == '.' {
				cpy1 := deepcopy(input)
				cpy1[x][y] = '#'
				isObstructed(cpy1, sx, sy, directions, &count)
			}
		}
	}
	fmt.Printf("Count: %v\n", count)

}

func findTiles(input [][]byte) {
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
			input[nx][ny] = 'X'
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
