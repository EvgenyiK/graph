package midleware

import (
    "testing"
)

var g ItemGraph

var graph1,_ = getGraph(6)
var graph2,_ = getGraph(8)
var graph3,_ = getGraph(9)

func fillGraph() {
    nA := Node{graph1.ID}
    nB := Node{graph2.ID}
    nC := Node{graph3.ID}
    
    g.AddNode(&nA)
    g.AddNode(&nB)
    g.AddNode(&nC)
    

    g.AddEdge(&nA, &nB)
    g.AddEdge(&nA, &nC)
    g.AddEdge(&nB, &nA)
    g.AddEdge(&nC, &nB)
    g.AddEdge(&nC, &nA)
}

func TestAdd(t *testing.T) {
    fillGraph()
    g.String()
}
