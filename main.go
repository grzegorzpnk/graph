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
			graph.addVertex(i, "MEC")
		} else {
			graph.addVertex(i, "CELL")
		}
	}

	graph.addEdge(1, 4)
	graph.addEdge(2, 5)
	graph.addEdge(3, 2)
	graph.addEdge(1, 0)
	graph.addEdge(4, 5)
	graph.printGraph()
}
