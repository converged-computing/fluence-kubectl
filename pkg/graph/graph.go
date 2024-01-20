package graph

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// AddCores adds cores to the node type
func (n *NodeType) AddCores(count int) {
	n.Cores += count
}

// AddMemory adds memory to the type
func (n *NodeType) AddMemory(mb int) {
	n.Memory += mb
}

// Uid returns a unique id based on cores and memory
func (n *NodeType) Uid() string {
	return fmt.Sprintf("%d-cores-%d-memory-mb", n.Cores, n.Memory)
}

// AddNode adds a node type to the graph summary
func (s *GraphSummary) AddNode(node *NodeType) {
	uid := node.Uid()
	_, ok := s.Nodes[uid]

	// This adds with correct memory / cores, just missing count
	if !ok {
		s.Nodes[uid] = node
	}

	meta := s.Nodes[uid]
	meta.Count += 1
	s.Nodes[uid] = meta
}

// Summary parses over the graph and returns a summary of unique nodes
// This is a quick and dirty parsing, and likely can be improved
func (g *Graph) Summary() {

	// Organize entire graph by id, and ensure we have a NodeType object for
	// each, where we will sum cores/memory across the graph
	nodes := map[string]Node{}
	nodeTypes := map[string]*NodeType{}
	for _, node := range g.Nodes {
		nodes[node.ID] = node
	}

	// Now for each edge, look for core and memory and define for each unique node
	for _, edge := range g.Edges {
		source := nodes[edge.Source]
		target := nodes[edge.Target]

		// If we have a node, always add to nodeType
		if source.Metadata.Type == "node" {
			_, ok := nodeTypes[source.ID]
			if !ok {
				nodeTypes[source.ID] = &NodeType{}
			}
			meta := nodeTypes[source.ID]

			// We've found a core for a node
			if target.Metadata.Type == "core" {
				meta.AddCores(target.Metadata.Size)
			}

			// The same, but for memory
			if target.Metadata.Type == "memory" {
				meta.AddMemory(target.Metadata.Size)
			}

		}
	}

	// Go through this listing and determine unique combinations and count then
	summary := GraphSummary{Nodes: map[string]*NodeType{}}
	for _, meta := range nodeTypes {
		summary.AddNode(meta)
	}

	// Write out table with nodes
	t := table.NewWriter()
	t.SetTitle("Resource Graph Summary")
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"", "Nodes", "Cores/Node", "Memory/Node"})
	t.AppendSeparator()

	// Keep totals as we go
	totals := map[string]int{"nodes": 0, "cores": 0, "memory": 0}
	for _, node := range summary.Nodes {
		if node.Count != 0 {
			t.AppendRow([]interface{}{"", node.Count, node.Cores, node.Memory})
			totals["nodes"] += int(node.Count)
			totals["cores"] += node.Cores
			totals["memory"] += node.Memory
		}
	}
	t.AppendSeparator()
	t.AppendFooter(table.Row{"Total", totals["nodes"], totals["cores"], totals["memory"]})
	t.SetStyle(table.StyleColoredBlackOnBlueWhite)
	t.Render()

}
