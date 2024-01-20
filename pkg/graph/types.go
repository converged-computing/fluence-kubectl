package graph

// ResourceGraph and nested structs describe a Flux resource graph
type ResourceGraph struct {
	Graph Graph `json:"graph"`
}

type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Paths struct {
	Containment string `json:"containment"`
}

type Edge struct {
	Source   string       `json:"source"`
	Target   string       `json:"target"`
	Metadata EdgeMetadata `json:"metadata"`
}

type EdgeMetadata struct {
	Name Paths `json:"name"`
}

type NodeMetadata struct {
	Type      string `json:"type"`
	Basename  string `json:"basename"`
	Name      string `json:"name"`
	ID        int    `json:"id"`
	UniqID    int    `json:"uniq_id"`
	Rank      int    `json:"rank"`
	Exclusive bool   `json:"exclusive"`
	Unit      string `json:"unit"`
	Size      int    `json:"size"`
	Paths     Paths  `json:"paths"`
}

type Node struct {
	ID       string       `json:"id"`
	Metadata NodeMetadata `json:"metadata"`
}

// Summary describes unique node types
type GraphSummary struct {
	// A lookup of node to unique identifier, and count
	Nodes map[string]*NodeType
}

// A NodeType describes a node resources and count
type NodeType struct {
	Count  int32
	Memory int
	Cores  int
}
