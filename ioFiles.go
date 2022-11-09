package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Edge struct {
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

	// -- OLD VERSION --
	// text, err := ioutil.ReadFile(filename)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// textString := strings.Replace(string(text), "\r", "", -1)
	// fmt.Println("\n textString:\n", textString, "\n -- end --")
	// textSlice := strings.Split(string(textString), "\n")

	return lines
}

func WriteFile(filename string, text string) {
	err := os.WriteFile(filename, []byte(text), 0644)
	// check(err)

	// err := ioutil.WriteFile(filename, []byte(text), 0644)
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
edges: slice of edges (struct Edge)
*/
func readInitFile(filename string) (uint32, []uint32, []Edge) {
	reg := regexp.MustCompile(`(?P<node1>\d+) (?P<node2>\d+) (?P<weight>-?\d+\.?\d*)`)
	lines := ReadFile(filename)

	// get number of vertexes
	lenVertex64, err := strconv.ParseUint(strings.Trim(lines[0], " "), 10, 32)
	if err != nil {
		log.Print("Unable to convert string '", lines[0], "' to uint32; char count: ", len(lines[0]))
		log.Fatal(err)
	}
	lenVertex := uint32(lenVertex64 + 1) // add the 'null' vertex (0)
	fmt.Println("lenVertex: ", lenVertex)

	fmt.Println("get edges")
	// get edges
	var edges []Edge                                    // edges (order doesn't matter)
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
		edges = append(edges, Edge{uint32(node1id64), uint32(node2id64), weight}) // add edge to slice
	}

	return lenVertex, neighCount, edges

}

func tree2text(tree []*TreeNode) string {
	var text_bytes bytes.Buffer
	for i := 0; i < len(tree); i++ {
		line := fmt.Sprint(tree[i].id) + "\t\t" + fmt.Sprint(tree[i].father) + "\t\t" + fmt.Sprintf("%.2f", tree[i].cost) + "\n"
		text_bytes.WriteString(line)
	}
	return text_bytes.String()
}
