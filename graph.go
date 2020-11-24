package graph

// Node describes a node in the graph and its neighbours
type Node struct {
	key        string
	neighbours []Node
}

func (n *Node) addNeighbour(neighbour Node) {
	n.neighbours = append(n.neighbours, neighbour)
}

// Graph is a datastructer with Nodes & weighted edges
type Graph struct {
	nodes    map[string]Node
	edges    map[string]int
	directed bool
}

// NewGraph returns a new graph
func NewGraph(directed bool) *Graph {
	return &Graph{
		directed: true,
		nodes:    make(map[string]Node),
		edges:    make(map[string]int),
	}
}

// AddNode creates a new Node in the graph given a key
func (g *Graph) AddNode(key string) {
	// g.Nodes = append(g.Nodes, Node{key: key})
	// fmt.Printf("%+v\n", g.nodes["n"])
	g.nodes[key] = Node{
		key: key,
	}
}

// GetNode returns a Node given a key
func (g *Graph) GetNode(key string) Node {
	return g.nodes[key]
}
