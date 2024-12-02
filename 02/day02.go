package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func main() {
    f, err := os.Open("task.txt")
    //f, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    
    scanner := bufio.NewScanner(f)
    i := 1
    result1 := 0
    result2 := 0
    for scanner.Scan() {
        split := strings.Split(scanner.Text(), " ")
        result1 += isSafe(split)
        result2 += isDampenedSafe(split, false)

        if err != nil {
            panic(err)
        }
        i = i + 1
    }
    fmt.Printf("isSafe (1)%i\n", result1)
    fmt.Printf("isDampenedSafe (2)%i\n", result2)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
func isDampenedSafe(split []string, error bool) int {
    incr := true
    prevs, split2 := split[0], split[1:]
    prev, err := strconv.Atoi(prevs)
    if err != nil {
        panic(err)
    }
    for i, v := range split2 {
        next, err := strconv.Atoi(v)
        if err != nil {
            panic(err)
        }
        if( i == 0) {
            incr = prev < next 
        }
        if (next == prev || prev < next != incr || prev - next < -3 || prev - next > 3  ) {
            if (error) {
                return 0
            }
            psplit := append(make([]string, 0, len(split[:i])), split[:i]...)
            x := split[i+1:]
            psplit = append(psplit, x...)
            // remove the prev-Element (which is split2[i-1] aka split[i])
            if (isDampenedSafe(psplit, true) == 1 ) {
                return 1
            }
            nsplit := append(make([]string, 0, len(split[:i+1])), split[:i+1]...)
            y := split[i+2:]
            nsplit = append(nsplit, y...)
            // remove the next-Element
            if (i+2 <= len(split) && isDampenedSafe(nsplit, true) == 1 ){
                return 1
            }
            // remove first Element, maybe start was wrong?
            if (isDampenedSafe(split[1:], true) == 1) {
                return 1
            }
            // remove last Element, maybe end was wrong?
            if (isDampenedSafe(split[:len(split)-1], true) == 1) {
                return 1
            }
            
            return 0
        }
        prev = next
    }
    return 1
}

func isSafe(split []string) int {
    incr := true
    prevs, split2 := split[0], split[1:]
    prev, err := strconv.Atoi(prevs)
    if err != nil {
        panic(err)
    }
    for i, v := range split2 {
        next, err := strconv.Atoi(v)
        if err != nil {
            panic(err)
        }
        if( i == 0) {
            incr = prev < next 
        }
        if (next == prev || prev < next != incr || prev - next < -3 || prev - next > 3  ) {
            return 0
        }
        prev = next
    }
    return 1
}
