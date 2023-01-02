package main

import (
	"os"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func main() {
	g := graph.New(graph.IntHash, graph.Directed())

	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_ = g.AddVertex(3)
	_ = g.AddVertex(4)

	_ = g.AddEdge(1, 2)
	_ = g.AddEdge(1, 3)
	_ = g.AddEdge(3, 4)
	_ = g.AddEdge(1, 2, graph.EdgeAttribute("color", "red"))

	file, _ := os.Create("my-graph.gv")
	_ = draw.DOT(g, file)
}
