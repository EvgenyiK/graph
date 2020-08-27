package main

import (
	"fmt"
	"log"
	"net/http"


	m "github.com/EvgenyiK/graph/server/midleware"
	"github.com/EvgenyiK/graph/server/router"
)

func main() {
	r := router.Router()
	m.GraphPrint()
	fmt.Println("Starting server on the port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
