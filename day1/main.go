package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, val)
	}
	return lines, scanner.Err()

}
func sol1(lines []int) int{
  var greater []int
        for i := 0; i < len(lines); i++ {
                if (i + 1) < len(lines) {
                        if lines[i+1] > lines[i] {
                                greater = append(greater, lines[i+1])

                        }
                }
        }
        fmt.Println(len(greater))
return len(greater)
}
func sol2(lines []int) int {
  var greater []int
  s1:= 0
  s2 := 0
        for i := 0; i < len(lines); i++ {
                if (i + 3) < len(lines) {
                        s1 = lines[i] + lines [i+1] + lines [i+2]
			s2 = lines[i+1] + lines [i+2] + lines [i+3]
                        if s2 > s1 {
                                greater = append(greater, s2)

                        }
                }
        }
        fmt.Println(len(greater))
return len(greater)
}
func main() {
	lines, err := readFile("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
        sol1(lines)
        sol2(lines)
}
