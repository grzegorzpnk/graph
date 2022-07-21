package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

//sprawdz co sie stanie jak w obiekcie nie zaznaczysz adjacent list jako pole marshowalne
func getVertexesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	/*	for i, _ := range graph.Vertices {
		fmt.Fprintf(w, "Vertex: ")
		json.NewEncoder(w).Encode(graph.Vertices[i].Id)

		if len(graph.Vertices[i].Adjacent) != 0 {
			fmt.Fprintf(w, "Adjacent to: ")
			for j, _ := range graph.Vertices[i].Adjacent {
				fmt.Fprintf(w, strconv.Itoa(graph.Vertices[i].Adjacent[j].Id))
				fmt.Fprintf(w, ", ")
			}
			fmt.Fprintln(w, "\n")

		} else {
			fmt.Fprintf(w, "Adjacent to: nil\n\n")
		}
	}*/
	json.NewEncoder(w).Encode(graph.Vertices)
}

func getVertexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	/*	params := mux.Vars(r)

		for i, v := range graph.Vertices {
			if strconv.Itoa(v.Id) == params["Id"] {
				fmt.Fprintf(w, "Vertex: ")
				json.NewEncoder(w).Encode(graph.Vertices[i].Id)
				fmt.Fprintf(w, "Adjacent to: ")
				for j, _ := range graph.Vertices[i].Adjacent {
					fmt.Fprintf(w, strconv.Itoa(graph.Vertices[i].Adjacent[j].Id)+", ")
					//json.NewEncoder(w).Encode(graph.Vertices[i].Adjacent[j].Id)
				}
				return
			}
		}*/
	json.NewEncoder(w).Encode(&Graph{})

}
