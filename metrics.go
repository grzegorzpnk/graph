package main

type ClusterMetrics struct {
	CpuUsage    int `json:"cpuUsage"`
	MemoryUsage int `json:"memoryUsage"`
	RamUsage    int `json:"ramUsage"`
}

type NetworkMetrics struct {

	//ms
	Latency    float32 `json:"latency"`
	PacketDrop int     `json:"packetDrop"`
}
