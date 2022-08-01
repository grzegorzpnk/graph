package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var graph *Graph
var r mux.Router

func main() {

	graph = &Graph{}
	initializingGraph()

	r := mux.NewRouter()
	defineRouterHandlers()

	log.Fatal(http.ListenAndServe(":8000", r))
}

type Graph struct {
	Vertices []*Vertex
	Edges    []*Edge
}

type Vertex struct {
	Id         int    `json:"id"`
	Type       string `json:"type"` //MEC or CELL
	Neighbours []int  `json:"neighbours"`
}

type Edge struct {
	Source  int     `json:"source"`
	Target  int     `json:"target"`
	Latency float32 `json:"latency"`
}

func (g *Graph) getVertex(k int) *Vertex {
	//getVertexHandler return a pointer to the Vertex with a key int

	for i, v := range g.Vertices {
		if v.Id == k {
			return g.Vertices[i]
		}
	}
	return nil
}

func (g *Graph) addVertex(k int, kind string) {
	if containsVertex(g.Vertices, k) {
		err := fmt.Errorf("Vertex %v not added beacuse already exist vertex with the same id", k)
		fmt.Println(err.Error())
	} else {
		g.Vertices = append(g.Vertices, &Vertex{Id: k, Type: kind})
		fmt.Printf("added new vertex ID: %v", k)
	}
}

func (g *Graph) addEdge(from, to int, latency float32) {

	//get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v<-->%v)", from, to)
		fmt.Println(err.Error())
	} else if fromVertex.Type == toVertex.Type {
		err := fmt.Errorf("You cannot connect two Vertexes at the same type:  %v !\n", fromVertex.Type)
		fmt.Println(err.Error())
	} else if containsInt(fromVertex.Neighbours, to) || containsInt(toVertex.Neighbours, from) {
		err := fmt.Errorf("Edge between (%v--%v) already exist", from, to)
		fmt.Println(err.Error())
	} else {
		//add edge at vertexes instances
		fromVertex.Neighbours = append(fromVertex.Neighbours, to)
		toVertex.Neighbours = append(toVertex.Neighbours, from)

		//add edge at  Edges list
		g.Edges = append(g.Edges, &Edge{from, to, latency})
		fmt.Printf("New Edge added : %v --- %v \n", from, to)
	}
}

func (g *Graph) printGraph() {

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
		fmt.Printf("Edge between: %v and %v\n", v.Source, v.Target)
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

func defineRouterHandlers() {
	r.HandleFunc("/graph/vertex", getAllVertexesHandler).Methods("GET")
	r.HandleFunc("/graph/vertex/{Id}", getVertexHandler).Methods("GET")
	r.HandleFunc("/graph/vertex", createVertex).Methods("POST")

	r.HandleFunc("/graph/edge", getEdgesHandler).Methods("GET")
	//r.HandleFunc("/graph/edge/{Id}", getEdgeHandler).Methods("GET")
	r.HandleFunc("/graph/edge", createEdgeHandler).Methods("POST")

}

func initializingGraph() {

	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			graph.addVertex(i, "MEC")
		} else {
			graph.addVertex(i, "CELL")
		}
	}

	graph.addEdge(1, 4, 0)
	graph.addEdge(2, 5, 0)
	graph.addEdge(3, 2, 0)
	graph.addEdge(1, 0, 0)
	graph.addEdge(4, 5, 0)
	graph.printGraph()
}
