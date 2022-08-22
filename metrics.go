package main

type ClusterMetrics struct {
	cpuUsage    int `json:"cpuUsage"`
	memoryUsage int `json:"memoryUsage"`
	ramUsage    int `json:"ramUsage"`
}

type NetworkMetrics struct {

	//ms
	latency    float32 `json:"latency"`
	packetDrop int     `json:"packetDrop"`
}
