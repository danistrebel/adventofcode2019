package day6

import "testing"

func TestToalOrbit(t *testing.T) {
	graph := buildOrbitGraph([]string{
		"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H",
		"D)I", "E)J", "J)K", "K)L"})

	calculated := totalOrbitSize(graph)
	if calculated != 42 {
		t.Error("Got", calculated, "expected 42")
	}
}
