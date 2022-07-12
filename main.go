package main

import (
	"fmt"
)

func main() {

	fmt.Println("First commit")

	graph := &Graph{}

	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}
	graph.AddVertex(2)
	graph.AddVertex(5)
	graph.AddVertex(6)
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
	if contains(g.vertices, k) {
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
func (g *Graph) addEdge(from, to int) {

	//get vertex
	//fromVertex := g.getVertex(from)
	//toVertex := g.getVertex(to)
	//check error
	//add edge

}

//getVertex return a pointer to the Vertex with a key int
func (g *Graph) getVertex(k int) *Vertex {

	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		} else {
			err := fmt.Errorf("cannot find Vertex with %v key", k)
			fmt.Println(err.Error())
		}
	}
	return nil
}
