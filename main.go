package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var graph *Graph

func main() {

	graph = &Graph{}

	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			graph.AddVertex(i, "MEC")
		} else {
			graph.AddVertex(i, "CELL")
		}
	}

	graph.addEdge(1, 3)
	graph.addEdge(2, 4)
	graph.addEdge(3, 1)
	graph.addEdge(1, 5)
	graph.addEdge(4, 5)
	graph.addEdge(4, 5)
	graph.Print()

	//create router
	r := mux.NewRouter()

	//create router handlers
	r.HandleFunc("/graph/vertex", getVertexesHandler).Methods("GET")
	r.HandleFunc("/graph/vertex/{Id}", getVertexHandler).Methods("GET")
	r.HandleFunc("/graph/vertex", createVertex).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

//graph structure
type Graph struct {
	Vertices []*Vertex
	Edges    []*Edge
}

type Edge struct {
	Source Vertex
	Target Vertex
}

//vertex structure
type Vertex struct {
	Id         int    `json:"id"`
	Kind       string `json:"type"` //MEC or CELL
	Neighbours []int  `json:"neighbours"`
}

func (g *Graph) Print() {

	//print vertexes
	for _, v := range g.Vertices {
		fmt.Printf("\nVertex: %v : ", v.Id)
		for _, v := range v.Neighbours {
			fmt.Printf("%v  ", v)
		}
	}
	fmt.Println()

	//print edges
	for _, v := range g.Edges {
		fmt.Printf("Edge between: %v and %v\n", v.Source.Id, v.Target.Id)
	}

}

//addVertex
func (g *Graph) AddVertex(k int, kind string) {
	if containsVertex(g.Vertices, k) {
		err := fmt.Errorf("Vertex %v not added beacuse already exist vertex with the same id", k)
		fmt.Println(err.Error())
	} else {
		g.Vertices = append(g.Vertices, &Vertex{Id: k, Kind: kind})
	}
}

func containsVertex(s []*Vertex, k int) bool {

	for _, v := range s {
		if k == v.Id {
			return true
		}
	}
	return false

}

func containsInt(s []int, k int) bool {

	for _, v := range s {
		if k == v {
			return true
		}
	}
	return false

}

//addEdge
func (g *Graph) addEdge(from, to int) {

	/*	//get vertex
		fromVertex := g.getVertex(from)
		toVertex := g.getVertex(to)
		//check error
		if fromVertex == nil || toVertex == nil {
			err := fmt.Errorf("Invalid edge (%v<-->%v)", from, to)
			fmt.Println(err.Error())
		} else if containsVertex(fromVertex.Adjacent, to) || containsVertex(toVertex.Adjacent, from) {
			err := fmt.Errorf("Edge between (%v--%v) already exist", from, to)
			fmt.Println(err.Error())
		} else {
			//add edge
			fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
			toVertex.Adjacent = append(toVertex.Adjacent, fromVertex)
		}*/

	/////SECOND IMPLEMENTATION/////

	//get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v<-->%v)", from, to)
		fmt.Println(err.Error())
	} else if containsInt(fromVertex.Neighbours, to) || containsInt(toVertex.Neighbours, from) {
		err := fmt.Errorf("Edge between (%v--%v) already exist", from, to)
		fmt.Println(err.Error())
	} else {
		//add edge
		fromVertex.Neighbours = append(fromVertex.Neighbours, to)
		toVertex.Neighbours = append(toVertex.Neighbours, from)

		g.Edges = append(g.Edges, &Edge{*g.getVertex(to), *g.getVertex(from)})
	}

}

//getVertexHandler return a pointer to the Vertex with a key int
func (g *Graph) getVertex(k int) *Vertex {

	for i, v := range g.Vertices {
		if v.Id == k {
			return g.Vertices[i]
		}
	}
	return nil
}
