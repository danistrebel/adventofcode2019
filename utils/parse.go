package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return lines
}

func ParseIntArrays(path string) [][]int {
	lines := ParseLines(path)
	var ints [][]int
	for _, line := range lines {
		var n []int
		numbers := strings.Split(line, ",")
		for _, number := range numbers {
			parsedInt, _ := strconv.Atoi(number)
			n = append(n, parsedInt)
		}
		ints = append(ints, n)
	}
	return ints
}
