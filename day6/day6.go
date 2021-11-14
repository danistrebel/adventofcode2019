package day6

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func buildOrbitGraph(links []string) map[string][]string {
	orbitGraph := make(map[string][]string)
	for _, l := range links {
		rel := strings.Split(l, ")")
		if prev, ok := orbitGraph[rel[0]]; ok {
			orbitGraph[rel[0]] = append(prev, rel[1])
		} else {
			orbitGraph[rel[0]] = []string{rel[1]}
		}
	}
	return orbitGraph
}

func totalOrbitSize(orbitGraph map[string][]string) int {
	return tallySizes(orbitGraph, "COM", 0)
}

func tallySizes(orbitGraph map[string][]string, orbitCenter string, hopsFromCenter int) int {
	sum := hopsFromCenter
	for _, orbitObj := range orbitGraph[orbitCenter] {
		sum += tallySizes(orbitGraph, orbitObj, hopsFromCenter+1)
	}
	return sum
}

func PrintSolution() {
	lines := utils.ParseLines("./inputs/day6.txt")
	graph := buildOrbitGraph(lines)
	totalSize := totalOrbitSize(graph)
	fmt.Println("Total Orbit Size (part 1)", totalSize)
}
