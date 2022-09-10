package main

// A slice with all the &vertexes
var AllVertex []*Vertex = make([]*Vertex, lenVertex+1) // 0 is empty

// type Vertexer interface {
// 	fillNeighbors([]Edge)
// 	createNeighbor(*Edge)
// 	deleteNeighbor(*Edge)
// }

// unidirectional edge (belongs to 1 node)
type Edge struct {
	neighbor *Vertex
	weight   float32
}

type Vertex struct {
	id    int // not necessary, remove after tests
	edges []Edge
}

// initial function, creates all vertexes
func createVertexes(lenVertex int) {
	for i := 0; i <= lenVertex; i++ {
		AllVertex[i] = &Vertex{i, []Edge{}}
	}
}

func (v *Vertex) fillNeighbors(nList []Edge) {
	v.edges = nList
}

func (v *Vertex) createEdge(edge Edge) {
	v.edges = append(v.edges, edge)
}

// not finished
func (v *Vertex) deleteNeighbor(neighborA *Edge) {
	for i, neighbor := range v.edges {
		if neighbor.neighbor == neighborA.neighbor {
			v.edges = append(v.edges[:i], v.edges[i+1:]...)
		}
	}
}

// create 2 edges with defined [weight] (one for each vertex)
func addEdge(node1 *Vertex, node2 *Vertex, weight float32) {
	node1.createEdge(Edge{node2, weight})
	node2.createEdge(Edge{node1, weight})
}
