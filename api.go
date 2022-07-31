package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func createVertex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vertex Vertex
	_ = json.NewDecoder(r.Body).Decode(&vertex)
	fmt.Println(vertex.Id)
	if containsVertex(graph.Vertices, vertex.Id) {
		err := fmt.Errorf("Vertex %v not added beacuse it is an existing key", vertex.Id)
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusConflict)
	} else {
		graph.AddVertex(vertex.Id)
		w.WriteHeader(http.StatusOK)
	}
}

func getVertexesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graph.Vertices)
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
