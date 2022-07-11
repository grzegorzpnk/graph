package main

import "fmt"

func main() {

	fmt.Println("First commit")

	graph := &Graph{}

	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}

	graph.Print()

}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex: %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v : ", v.key)
		}
	}
	fmt.Println()
}

//graph structure
type Graph struct {
	vertices []*Vertex
}

//vertex structure
type Vertex struct {
	key      int
	adjacent []*Vertex
}

//addVertex
func (g *Graph) AddVertex(k int) {
	if !contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added beacuse it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

func contains(s []*Vertex, k int) bool {

	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

//addEdge
