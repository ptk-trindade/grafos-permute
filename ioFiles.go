package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type rawEdge struct {
	vertex1 uint32
	vertex2 uint32
	weight  float64
}

// Reads a txt returning a slice of strings (separated by new line)
func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	br := bufio.NewReader(file)
	r, _, err := br.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	if r != '\uFEFF' {
		br.UnreadRune() // Not a BOM -- put the rune back
	}

	var lines []string
	scanner := bufio.NewScanner(br)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func WriteFile(filename string, text string) {
	err := os.WriteFile(filename, []byte(text), 0644)

	if err != nil {
		log.Fatal(err)
	}
}

/*
[initial function - insert the vertexes edges]
--- in:
filename: name of the file to be read
--- out:
lenVertex: number of vertexes
neighCount: number of neighbors of each vertex
edges: slice of edges (struct rawEdge)
*/
func readInitFile(filename string) (uint32, []uint32, []rawEdge) {
	reg := regexp.MustCompile(`(?P<node1>\d+) (?P<node2>\d+) (?P<weight>-?\d+\.?\d*)`)
	lines := ReadFile(filename)

	// get number of vertexes
	lenVertex64, err := strconv.ParseUint(strings.Trim(lines[0], " "), 10, 32)
	if err != nil {
		log.Print("Unable to convert string '", lines[0], "' to uint32; char count: ", len(lines[0]))
		log.Fatal(err)
	}
	lenVertex := uint32(lenVertex64 + 1) // add the 'null' vertex (0)

	// get edges
	var edges []rawEdge                                 // edges (order doesn't matter)
	var neighCount []uint32 = make([]uint32, lenVertex) // number of neighbors for each vertex
	for i := 1; i < len(lines); i++ {
		trimmed := strings.Trim(lines[i], " ")

		if !reg.MatchString(trimmed) {
			log.Fatal("ERROR readFile GetEdges - unable to read weights of line", i, ": ", trimmed)
		}

		var edgeSlice []string = reg.FindStringSubmatch(trimmed)

		node1id64, _ := strconv.ParseUint(edgeSlice[1], 10, 32)
		node2id64, _ := strconv.ParseUint(edgeSlice[2], 10, 32)
		weight, _ := strconv.ParseFloat(edgeSlice[3], 32)

		neighCount[node1id64]++
		neighCount[node2id64]++
		edges = append(edges, rawEdge{uint32(node1id64), uint32(node2id64), weight}) // add edge to slice
	}

	return lenVertex, neighCount, edges

}
