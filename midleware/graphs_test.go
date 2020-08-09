package midleware

import (
	"log"
    "testing"
)


var g ItemGraph

func fillGraph() {
    
    graph1,err:=getGraph(1)
    if err != nil {
		log.Fatalf("Error sqltatement")
	}
    graph2,err:=getGraph(2)
    if err != nil {
		log.Fatalf("Error sqltatement")
	}
    graph3,err:=getGraph(3)
    if err != nil {
		log.Fatalf("Error sqltatement")
	}

    nA := Node{graph1.Node}
    nB := Node{graph2.Node}
    nC := Node{graph3.Node}
    
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
