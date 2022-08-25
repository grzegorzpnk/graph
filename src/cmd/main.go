package main

import (
	"graph/src/api"
	"graph/src/pkg/topology"
	"log"
	"net/http"
)

//var graph *topology.Graph

func main() {

	var graph topology.Graph
	graph.InitializeGraph()

	httpRouter := api.NewRouter(&graph)
	log.Fatal(http.ListenAndServe("localhost:8080", httpRouter))

}
