package main

// Node represents a single node in a graph
type Node struct {
	Value    int
	Visited  bool
	Outgoing []*Edge
	Incoming []*Edge
	NumPaths int64
}

// NewNode returns a pointer to a new Node instance, with the Visited flag set to false, and
// an empty array of Edges
func NewNode(val int) *Node {
	return &Node{
		Value:    val,
		Visited:  false,
		NumPaths: int64(-1),
		Outgoing: make([]*Edge, 0),
		Incoming: make([]*Edge, 0),
	}
}

// LinkNode adds an Edge between the node this function is called on, and the given node
func (s *Node) LinkNode(n *Node) {
	e := &Edge{
		Source:      s,
		Destination: n,
	}

	s.Outgoing = append(s.Outgoing, e)
	n.Incoming = append(n.Incoming, e)
}
