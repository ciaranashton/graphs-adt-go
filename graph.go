package graph

import (
	"fmt"
	"sort"
)

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
		directed: directed,
		nodes:    make(map[string]Node),
		edges:    make(map[string]int),
	}
}

// AddNode creates a new Node in the graph given a key
func (g *Graph) AddNode(key string) {
	g.nodes[key] = Node{
		key: key,
	}
}

// GetNode returns a Node given a key
func (g *Graph) GetNode(key string) Node {
	return g.nodes[key]
}

// AddEdge creates a new edge between two nodes
func (g *Graph) AddEdge(key1 string, key2 string, weight int) {
	startNode := g.GetNode(key1)
	endNode := g.GetNode(key2)

	if g.directed {
		key := fmt.Sprint(key1, "-", key2)

		g.edges[key] = weight
		startNode.addNeighbour(endNode)
	} else {
		keys := []string{key1, key2}
		sort.Strings(keys)
		key := fmt.Sprint(keys[0], "-", keys[1])

		g.edges[key] = weight
		startNode.addNeighbour(endNode)
		endNode.addNeighbour(startNode)
	}
}

// GetEdge returns the weight of an edge
func (g *Graph) GetEdge(key1 string, key2 string) int {
	if g.directed {
		key := fmt.Sprint(key1, "-", key2)

		return g.edges[key]
	}

	keys := []string{key1, key2}
	sort.Strings(keys)
	key := fmt.Sprint(keys[0], "-", keys[1])

	return g.edges[key]
}
