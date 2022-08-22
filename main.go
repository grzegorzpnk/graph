package main

var graph *Graph

func main() {

	graph = &Graph{}
	initializingGraph()

	startServer()

}

func initializingGraph() {

	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			graph.addVertex(Vertex{Id: i, Type: "MEC", VertexMetrics: ClusterMetrics{20, 50, 80}})
		} else {
			graph.addVertex(Vertex{Id: i, Type: "CELL"})
		}
	}

	graph.addEdge(1, 4)
	graph.addEdge(2, 5)
	graph.addEdge(3, 2)
	graph.addEdge(1, 0)
	graph.addEdge(4, 5)
	graph.printGraph()
}
