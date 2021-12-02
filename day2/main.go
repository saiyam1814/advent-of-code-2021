package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	return lines, scanner.Err()

}
func sol1(lines []string) int {
	x := 0
	y := 0

	for i := range lines {

		if strings.Contains(lines[i], "forward") {
			split := strings.Split(lines[i], " ")
			val, _ := strconv.Atoi(split[1])
			x = x + val
		}
		if strings.Contains(lines[i], "down") {
			split := strings.Split(lines[i], " ")
			val, _ := strconv.Atoi(split[1])
			y = (y + val)
		}
		if strings.Contains(lines[i], "up") {
			split := strings.Split(lines[i], " ")
			val, _ := strconv.Atoi(split[1])
			y = (y - val)
		}
	}

	return (x * y)
}

func sol2(lines []string) int {
	x := 0
	y := 0
	aim := 0
	for i := range lines {

		if strings.Contains(lines[i], "forward") {
			split := strings.Split(lines[i], " ")
			val, _ := strconv.Atoi(split[1])
			if aim == 0 {
				x = x + val
			} else {
				x = x + val
				y = y + (val * aim)
			}
		}
		if strings.Contains(lines[i], "down") {
			split := strings.Split(lines[i], " ")
			val, _ := strconv.Atoi(split[1])
			aim = (aim + val)
		}
		if strings.Contains(lines[i], "up") {
			split := strings.Split(lines[i], " ")
			val, _ := strconv.Atoi(split[1])
			aim = (aim - val)
		}

	}
	return (x * y)

}

func main() {
	lines, err := readFile("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println(sol1(lines))
	fmt.Println(sol2(lines))

}
