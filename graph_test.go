package graph

import (
	"reflect"
	"testing"
)

func TestAddNodeToGraph(t *testing.T) {
	g1 := NewGraph(true)

	// fmt.Printf("%+v\n", g1)
	g1.AddNode("a")
	g1.AddNode("b")

	n1 := Node{key: "a"}
	n2 := Node{key: "b"}

	if !reflect.DeepEqual(g1.GetNode("a"), n1) {
		t.Fatalf(`Result: %+v`, g1)
	}
	if !reflect.DeepEqual(g1.GetNode("b"), n2) {
		t.Fatalf(`Result: %+v`, g1)
	}
}

func TestAddNeighbourToNode(t *testing.T) {
	n1 := Node{
		key: "a",
	}

	n1.addNeighbour(Node{key: "b"})
	n1.addNeighbour(Node{key: "c"})

	expected := Node{
		key: "a",
		neighbours: []Node{
			Node{
				key: "b",
			},
			Node{
				key: "c",
			},
		},
	}

	if !reflect.DeepEqual(n1, expected) {
		t.Fatalf(`Result: %+v, Expected: %+v`, n1, expected)
	}
}
