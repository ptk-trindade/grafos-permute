package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Reads a txt returning a slice of strings (separated by new line)
func ReadFile(filename string) []string {
	text, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	textString := string(text)

	textSlice := strings.Split(string(textString), "\n")

	return textSlice
}

// initial function - insert the vertexes edges
func populateEdges(vertexesAdd []*Vertex) {
	reg := regexp.MustCompile(`(?P<node1>\d+) - (?P<node2>\d+)\s+(?P<weight>\d+\.\d+)`)
	lines := ReadFile("MST.txt")

	for i := 1; i < len(lines)-1; i++ {
		trimmed := strings.Trim(lines[i], " ")

		if !reg.MatchString(trimmed) {
			log.Fatal("ERROR readFile GetEdges - unable to read weights of line", i, ": ", trimmed)
		}

		edgeSlice := reg.FindStringSubmatch(trimmed)

		node1id, _ := strconv.Atoi(edgeSlice[1])
		node2id, _ := strconv.Atoi(edgeSlice[2])
		weight, _ := strconv.ParseFloat(edgeSlice[3], 32)

		addEdge(vertexesAdd[node1id], vertexesAdd[node2id], float32(weight))

		// fmt.Printf("%#v\n", reg.FindStringSubmatch(trimmed))
	}

}
