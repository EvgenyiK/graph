package router

import(
	"github.com/EvgenyiK/graph/server/midleware"
	"github.com/gorilla/mux"
)

//Router маршрутизация
func Router() *mux.Router {
	router:= mux.NewRouter()

	router.HandleFunc("/api/graph/{id}", midleware.GetGraph).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/graph", midleware.GetAllGraph).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newgraph", midleware.CreateGraph).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/graph/{id}", midleware.UpdateGraph).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deletegraph/{id}", midleware.DeleteGraph).Methods("DELETE", "OPTIONS")

	return router
}