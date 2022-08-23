package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var r *mux.Router

func startServer() {

	r = mux.NewRouter()
	defineRouterHanlders()
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}

func defineRouterHanlders() {

	r.HandleFunc("/graph/vertex", getAllVertexesHandler).Methods("GET")
	r.HandleFunc("/graph/vertex/{Id}", getVertexHandler).Methods("GET")
	r.HandleFunc("/graph/vertex/{Id}/metrics", updateClusterMetrics).Methods("PUT")
	r.HandleFunc("/graph/vertex", createVertex).Methods("POST")

	r.HandleFunc("/graph/edge", getEdgesHandler).Methods("GET")
	r.HandleFunc("/graph/edge", createEdgeHandler).Methods("POST")
	r.HandleFunc("/graph/edge/{IdSource}/{IdTarget}/metrics", updateEdgeMetrics).Methods("POST")

}

func createEdgeHandler(w http.ResponseWriter, r *http.Request) {
	//todo: validate body of REST POST
	w.Header().Set("Content-Type", "application/json")

	var edge Edge
	_ = json.NewDecoder(r.Body).Decode(&edge)
	fmt.Printf("Client tries to add new Edge: %v --- %v \n", edge.Source, edge.Target)
	graph.addEdge(edge)

}

func updateClusterMetrics(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["Id"])
	var clusterMetrics ClusterMetrics
	_ = json.NewDecoder(r.Body).Decode(&clusterMetrics)

	if containsVertex(graph.Vertices, id) {
		graph.getVertex(id).VertexMetrics.updateClusterMetrics(clusterMetrics)
		w.WriteHeader(http.StatusOK)
		fmt.Printf("Client updates cluster metrics for vertex ID: %v\n", params["Id"])
	} else {
		err := fmt.Errorf("Vertex %v not updated beacuse it's not exist", id)
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusConflict)
	}
}

func updateEdgeMetrics(w http.ResponseWriter, r *http.Request) {

	/*	w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		idSource, _ := strconv.Atoi(params["IdSource"])
		idTarget, _ := strconv.Atoi(params["IdTarget"])
		var edgeMetrics NetworkMetrics
		_ = json.NewDecoder(r.Body).Decode(&edgeMetrics)
	*/

	//sprawdz czy istnieje dany link i go pobierz
	//update danych na łaczu
	/*if exist(graph.Vertices, id) {
		graph.getVertex(id).VertexMetrics.updateClusterMetrics(clusterMetrics)
		w.WriteHeader(http.StatusOK)
		fmt.Printf("Client updates cluster metrics for vertex ID: %v\n", params["Id"])
	} else {
		err := fmt.Errorf("Vertex %v not updated beacuse it's not exist", id)
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusConflict)
	}*/
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
		if containsAnyEdge(vertex) {
			vertex.Neighbours = nil
		}
		graph.addVertex(vertex)
		w.WriteHeader(http.StatusOK)
	}
}

func containsAnyEdge(vertex Vertex) bool {

	if vertex.Neighbours != nil {
		return true
	} else {
		return false
	}

}

func getEdgesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	for i, _ := range graph.Edges {
		json.NewEncoder(w).Encode(graph.Edges[i])
	}

}

func getAllVertexesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graph.Vertices)
}
