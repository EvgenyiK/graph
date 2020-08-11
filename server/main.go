package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EvgenyiK/graph/server/midleware"
	"github.com/EvgenyiK/graph/server/router"
)

func main() {
	r := router.Router()
	midleware.TestAdd()
	fmt.Println("Starting server on the port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
