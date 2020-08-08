package models

//GraphNode ...
type GraphNode struct {
    ID        int64 `json:"id"`
    Neighbors int64 `json:"neighboors"`
    Roots     int64 `json:"roots"`
}