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
		k := rel[0]
		v := rel[1]
		if prev, ok := orbitGraph[k]; ok {
			orbitGraph[k] = append(prev, v)
		} else {
			orbitGraph[k] = []string{v}
		}
	}
	return orbitGraph
}

func buildReverseOrbitGraph(links []string) map[string]string {
	orbitGraph := make(map[string]string)
	for _, l := range links {
		rel := strings.Split(l, ")")
		orbitGraph[rel[1]] = rel[0]
	}
	return orbitGraph
}

func totalOrbitSize(orbitGraph map[string][]string) int {
	return tallySizes(orbitGraph, "COM", 0)
}

func transfersCount(reverseGraph map[string]string) int {
	yPath := rootPath(reverseGraph, "YOU")
	sPath := rootPath(reverseGraph, "SAN")

	prefixLength := 0
	for {
		if prefixLength < len(yPath) && prefixLength < len(sPath) && yPath[prefixLength] == sPath[prefixLength] {
			prefixLength++
		} else {
			break
		}
	}

	return len(sPath) - prefixLength + len(yPath) - prefixLength - 2
}

func rootPath(reverseGraph map[string]string, leaf string) []string {
	path := []string{}

	curr := leaf
	for {
		path = append([]string{curr}, path...)
		parent, ok := reverseGraph[curr]
		if !ok {
			break
		} else {
			curr = parent
		}
	}
	return path
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
	reverseGraph := buildReverseOrbitGraph(lines)
	transferCounts := transfersCount(reverseGraph)
	fmt.Println("TransferCounts (part 2)", transferCounts)

}
