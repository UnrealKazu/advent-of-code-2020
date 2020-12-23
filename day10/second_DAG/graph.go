package main

import "sort"

// JoltGraph is a DAG representing the adapters
type JoltGraph struct {
	Root *Node
	Sink *Node
}

// New creates a new JoltGraph based on the given list of adapter values
// This given list must contain the outlet and device values, but does not necessarily have to be sorted
func New(values []int) *JoltGraph {
	sort.Ints(values)

	// first, create all nodes
	nodes := make([]*Node, len(values))
	for i, j := range values {
		nodes[i] = NewNode(j)
	}

	for i, n := range nodes {
		for j := i + 1; j <= i+3; j++ {
			if j > len(nodes)-1 {
				break
			}

			// check the next three nodes for their diff
			neigh := nodes[j]

			if (neigh.Value - n.Value) <= 3 {
				// a possible route, so connect it to n
				n.LinkNode(neigh)
			}
		}
	}

	return &JoltGraph{
		Root: nodes[0],
		Sink: nodes[len(nodes)-1],
	}
}

// GetNrofAllPaths returns the number of all possible paths
func (g *JoltGraph) GetNrofAllPaths() int64 {
	// start at the sink, and work our way back
	return traverseAllPaths(g.Root)
}

func traverseAllPaths(n *Node) int64 {
	if len(n.Outgoing) == 0 {
		return int64(1)
	}

	if n.NumPaths == -1 {
		sum := int64(0)

		for _, e := range n.Outgoing {
			sum += traverseAllPaths(e.Destination)
		}

		n.NumPaths = sum
	}

	return n.NumPaths
}
