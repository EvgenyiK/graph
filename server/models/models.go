package models

//GraphNode ...
type GraphNode struct {
	ID   int64 `json:"id"`
	Node string `json:"node"`
	Weight int64 `json:"weight"`
}
