package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func createEdgeHandler(w http.ResponseWriter, r *http.Request) {
	//todo: validate body of REST POST
	w.Header().Set("Content-Type", "application/json")

	var edge Edge
	_ = json.NewDecoder(r.Body).Decode(&edge)
	fmt.Printf("Client tries to add new Edge: %v --- %v \n", edge.Source, edge.Target)
	graph.addEdge(edge.Source, edge.Target)

}

func defineRouterHandlers() {

	r.HandleFunc("/graph/vertex", getAllVertexesHandler).Methods("GET")
	r.HandleFunc("/graph/vertex/{Id}", getVertexHandler).Methods("GET")
	r.HandleFunc("/graph/vertex", createVertex).Methods("POST")

	r.HandleFunc("/graph/edge", getEdgesHandler).Methods("GET")
	//r.HandleFunc("/graph/edge/{Id}", getEdgeHandler).Methods("GET")
	r.HandleFunc("/graph/edge", createEdgeHandler).Methods("POST")

}

func getVertexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, v := range graph.Vertices {
		if strconv.Itoa(v.Id) == params["Id"] {
			json.NewEncoder(w).Encode(graph.Vertices[i])
		}
	}

}

func getAllVertexesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graph.Vertices)
}

func createVertex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var vertex Vertex
	_ = json.NewDecoder(r.Body).Decode(&vertex)
	fmt.Printf("Client tries to add new vertex ID: %v\n", vertex.Id)
	if containsVertex(graph.Vertices, vertex.Id) {
		err := fmt.Errorf("Vertex %v not added beacuse it is an existing key", vertex.Id)
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusConflict)
	} else {
		graph.addVertex(vertex.Id, vertex.Type)
		w.WriteHeader(http.StatusOK)
	}
}

func getEdgesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	for i, _ := range graph.Edges {
		json.NewEncoder(w).Encode(graph.Edges[i])
	}

}
