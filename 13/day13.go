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
	findFewestTokens(input, 100, 100, 0)
	findFewestTokens(input, 10000000000000, 10000000000000, 10000000000000)

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

type tupel struct {
	x, y int
}

func solveRulesBetter(adepth int, bdepth int, tokens []tupel, solution tupel, smincount int) int {

	det := tokens[0].x*tokens[1].y - tokens[1].x*tokens[0].y
	adet := solution.x*tokens[1].y - solution.y*tokens[1].x
	bdet := solution.y*tokens[0].x - solution.x*tokens[0].y

	if det == 0 {
		return smincount
	}
	if adet%det != 0 || bdet%det != 0 {
		return smincount
	}
	a := adet / det
	b := bdet / det
	if a > adepth || b > bdepth || a < 0 || b < 0 {
		return smincount
	}
	return a*3 + b
}

func findFewestTokens(input [][]byte, adepth int, bdepth int, offset int) {
	defer timerHighRes("Day 13 Part X solveRules")()

	var tokens []tupel
	var solution tupel

	smincount := adepth*3 + bdepth + 1

	count := 0
	for _, line := range input {
		string := string(line)
		if strings.Contains(string, "Button") {
			a, _ := strconv.Atoi(strings.Split(strings.Split(string, "+")[1], ",")[0])
			b, _ := strconv.Atoi(strings.Split(string, "+")[2])
			tokens = append(tokens, tupel{a, b})
		}
		if strings.Contains(string, "Prize") {
			a, _ := strconv.Atoi(strings.Split(strings.Split(string, "=")[1], ",")[0])
			b, _ := strconv.Atoi(strings.Split(string, "=")[2])
			solution = tupel{a + offset, b + offset}

			pcount := solveRulesBetter(adepth, bdepth, tokens, solution, smincount)

			if pcount < smincount {
				count += pcount
			}
			tokens = make([]tupel, 0)

		}

	}
	fmt.Printf("Count: %v\n", count)
}
