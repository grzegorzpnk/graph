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

	graph.addEdge(Edge{1, 4, NetworkMetrics{1.3, 10}})
	graph.addEdge(Edge{2, 5, NetworkMetrics{1.3, 10}})
	graph.addEdge(Edge{3, 2, NetworkMetrics{1.3, 10}})
	graph.addEdge(Edge{1, 0, NetworkMetrics{1.3, 10}})
	graph.addEdge(Edge{4, 5, NetworkMetrics{1.3, 10}})
	graph.printGraph()
}
