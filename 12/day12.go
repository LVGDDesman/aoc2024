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

	findPrice(input)

	findBulkPrice(input)

}

type dirs struct {
	dx, dy int
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

func findPrice(input [][]byte) {
	defer timerHighRes("Day 12 Part 1 findPrice")()
	count := 0
	x := 0
	y := 0
	visited := make(map[int]map[int]bool)
	for {
		x, y = findNextArea(input, x, y, visited)
		if x == -1 {
			break
		}
		//fmt.Printf("Starting at %v %v with type %v ", x, y, string(input[x][y]))
		area, perimeter := getAreaAndPerimeter(input, x, y, visited)
		//fmt.Printf("Resulting in %v %v (%v)\n", area, perimeter, area*perimeter)

		count += area * perimeter

	}
	fmt.Printf("Count: %v\n", count)
}

func findBulkPrice(input [][]byte) {
	defer timerHighRes("Day 12 Part 2 findBulkPrice")()
	count := 0
	x := 0
	y := 0
	visited := make(map[int]map[int]bool)
	for {
		x, y = findNextArea(input, x, y, visited)
		if x == -1 {
			break
		}
		area, sides := getAreaAndSides(input, x, y, visited, make(map[int]map[int]byte))
		// fmt.Printf("Starting at %v %v with type %v ", x, y, string(input[x][y]))
		// fmt.Printf("%v %v (%v)\n", area, sides, area*sides)

		count += area * sides

	}
	fmt.Printf("Count: %v\n", count)
}

func getAreaAndSides(input [][]byte, x int, y int, visited map[int]map[int]bool, sides map[int]map[int]byte) (int, int) {
	area := 1
	csides := 0
	ptype := input[x][y]
	if visited[x] == nil {
		visited[x] = make(map[int]bool)
	}
	visited[x][y] = true

	directions := []dirs{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for i, dir := range directions {
		nx := x + dir.dx
		ny := y + dir.dy
		if nx < 0 || nx >= len(input) || ny < 0 || ny >= len(input[nx]) || input[nx][ny] != ptype {
			// +1 since nx/ny could be -1 -> offset everything
			if sides[nx+1] == nil {
				sides[nx+1] = make(map[int]byte)
			}
			sides[nx+1][ny+1] = sides[nx+1][ny+1] | 2<<i

			prev := directions[(i+1)%4]
			next := directions[(i+3)%4]
			if sides[nx+1+prev.dx] == nil || (sides[nx+1+prev.dx][ny+1+prev.dy])&(2<<i) == 0 {
				if sides[nx+1+next.dx] == nil || (sides[nx+1+next.dx][ny+1+next.dy])&(2<<i) == 0 {
					//fmt.Printf("    %v %v new Side Found: %v, %v, (tested(%v, %v : %v) and (%v, %v: %v)) == %v \n", x, y, nx, ny, nx+prev.dx, ny+prev.dy, sides[nx+1+prev.dx][ny+1+prev.dy], nx+next.dx, ny+next.dy, sides[nx+1+next.dx][ny+1+next.dy], 2<<i)
					csides += 1
				}
			}
			if sides[nx+1+prev.dx] != nil && sides[nx+1+next.dx] != nil && sides[nx+1+prev.dx][ny+1+prev.dy]&(2<<i) == 2<<i && sides[nx+1+next.dx][ny+1+next.dy]&(2<<i) == 2<<i {
				//correction if something like x _ x in sides happens
				//fmt.Printf("Correcting at %v, %v", x, y)
				csides -= 1
			}
		} else {
			if input[nx][ny] == ptype && (visited[nx] == nil || !visited[nx][ny]) {
				narea, ncsides := getAreaAndSides(input, nx, ny, visited, sides)
				area += narea
				csides += ncsides
			}
		}
	}
	return area, csides
}

// func getArea(input [][]byte, x int, y int, visited map[int]map[int]bool) int {
// 	area := 1
// 	ptype := input[x][y]
// 	if visited[x] == nil {
// 		visited[x] = make(map[int]bool)
// 	}
// 	visited[x][y] = true

// 	directions := []dirs{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
// 	for _, dir := range directions {
// 		nx := x + dir.dx
// 		ny := y + dir.dy
// 		if nx < 0 || nx >= len(input) || ny < 0 || ny >= len(input[nx]) || input[nx][ny] != ptype {
// 		} else {
// 			if input[nx][ny] == ptype && (visited[nx] == nil || !visited[nx][ny]) {
// 				area += getArea(input, nx, ny, visited)
// 			}
// 		}

// 	}
// 	return area
// }

func getAreaAndPerimeter(input [][]byte, x int, y int, visited map[int]map[int]bool) (int, int) {
	area := 1
	perimeter := 0
	ptype := input[x][y]
	if visited[x] == nil {
		visited[x] = make(map[int]bool)
	}
	visited[x][y] = true

	directions := []dirs{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, dir := range directions {
		nx := x + dir.dx
		ny := y + dir.dy
		if nx < 0 || nx >= len(input) || ny < 0 || ny >= len(input[nx]) || input[nx][ny] != ptype {
			perimeter += 1
			continue
		} else {
			if input[nx][ny] == ptype && (visited[nx] == nil || !visited[nx][ny]) {
				narea, nperimeter := getAreaAndPerimeter(input, nx, ny, visited)
				area += narea
				perimeter += nperimeter
			}
		}

	}
	return area, perimeter
}

func findNextArea(input [][]byte, x int, y int, visited map[int]map[int]bool) (int, int) {
	for {
		if visited[x] == nil {
			return x, y
		}
		if visited[x][y] != true {
			return x, y
		}
		y += 1
		if y == len(input[x]) {
			y = 0
			x += 1
		}
		if x == len(input) {
			return -1, -1
		}
	}
}
