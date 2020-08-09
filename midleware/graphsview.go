package midleware

import (
	"log"
    //"testing"
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
    graph4,err:=getGraph(4)
    if err != nil {
		log.Fatalf("Error sqltatement")
    }
    graph5,err:=getGraph(5)
    if err != nil {
		log.Fatalf("Error sqltatement")
    }
    graph6,err:=getGraph(6)
    if err != nil {
		log.Fatalf("Error sqltatement")
	}

    nA := Node{graph1.Node}
    nB := Node{graph2.Node}
    nC := Node{graph3.Node}
    nD := Node{graph4.Node}
    nE := Node{graph5.Node}
    nF := Node{graph6.Node}
    
    g.AddNode(&nA)
    g.AddNode(&nB)
    g.AddNode(&nC)
    g.AddNode(&nD)
    g.AddNode(&nE)
    g.AddNode(&nF)
    

    g.AddEdge(&nA, &nB)
    g.AddEdge(&nA, &nC)
    g.AddEdge(&nB, &nE)
    g.AddEdge(&nC, &nE)
    g.AddEdge(&nE, &nF)
    g.AddEdge(&nD, &nA)
}

func TestAdd() {
    fillGraph()
    g.String()
}
