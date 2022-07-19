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
	/*
		var vertex Vertex
		_ = json.NewDecoder(r.Body)

	*/
}

func getVertexesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	for i, _ := range graph.Vertices {
		fmt.Fprintf(w, "Vertex: ")
		json.NewEncoder(w).Encode(graph.Vertices[i].Key)

		if len(graph.Vertices[i].Adjacent) != 0 {
			fmt.Fprintf(w, "Adjacent to: ")
			for j, _ := range graph.Vertices[i].Adjacent {
				fmt.Fprintf(w, strconv.Itoa(graph.Vertices[i].Adjacent[j].Key))
				fmt.Fprintf(w, ", ")
			}
			fmt.Fprintln(w, "\n")

		} else {
			fmt.Fprintf(w, "Adjacent to: nil\n\n")
		}
	}
}

func getVertexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, v := range graph.Vertices {
		if strconv.Itoa(v.Key) == params["id"] {
			fmt.Fprintf(w, "Vertex: ")
			json.NewEncoder(w).Encode(graph.Vertices[i].Key)
			fmt.Fprintf(w, "Adjacent to: ")
			for j, _ := range graph.Vertices[i].Adjacent {
				fmt.Fprintf(w, strconv.Itoa(graph.Vertices[i].Adjacent[j].Key)+", ")
				//json.NewEncoder(w).Encode(graph.Vertices[i].Adjacent[j].Key)
			}
			return
		}
	}
	json.NewEncoder(w).Encode(&Graph{})

}
