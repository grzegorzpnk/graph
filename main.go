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
		graph.AddVertex(i)
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
	r.HandleFunc("/graph/vertex/{id}", getVertexHandler).Methods("GET")
	r.HandleFunc("graph/vertex/", createVertex).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

//graph structure
type Graph struct {
	Vertices []*Vertex `json:"Verticies"`
}

//vertex structure
type Vertex struct {
	Key int `json:"Key"`
	//Adj 	 []int 	   `json:"Adj"`
	Adjacent []*Vertex `json:"Adjacent"`
}

func (g *Graph) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("\nVertex: %v : ", v.Key)
		for _, v := range v.Adjacent {
			fmt.Printf("%v  ", v.Key)
		}
	}
	fmt.Println()
}

//addVertex
func (g *Graph) AddVertex(k int) {
	if contains(g.Vertices, k) {
		err := fmt.Errorf("Vertex %v not added beacuse it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.Vertices = append(g.Vertices, &Vertex{Key: k})
	}
}

func contains(s []*Vertex, k int) bool {

	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false
}

//addEdge
func (g *Graph) addEdge(from, to int) {

	//get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v<-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.Adjacent, to) || contains(toVertex.Adjacent, from) {
		err := fmt.Errorf("Edge between (%v--%v) already exist", from, to)
		fmt.Println(err.Error())
	} else {
		//add edge
		fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
		toVertex.Adjacent = append(toVertex.Adjacent, fromVertex)
	}
}

//getVertexHandler return a pointer to the Vertex with a key int
func (g *Graph) getVertex(k int) *Vertex {

	for i, v := range g.Vertices {
		if v.Key == k {
			return g.Vertices[i]
		}
	}
	return nil
}
