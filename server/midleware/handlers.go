package midleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	//"strings"

	"github.com/EvgenyiK/graph/server/models"

	"github.com/gorilla/mux"   //пакет для роутов
	"github.com/joho/godotenv" //пакет для чтения конфига
	_ "github.com/lib/pq"      //драйвер для работы с postgresql
)

//формат ответа
type response struct {
	ID      int64  `json:"id.omitempty"`
	Message string `json:"message,omitempty"`
}

//подключение к postresql
func createConnection() *sql.DB {
	//загрузка конфига
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//открытие соединения
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	//проверка соединения
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	//возврат соединения
	return db
}

//CreateGraph Создание графа
func CreateGraph(w http.ResponseWriter, r *http.Request) {
	//настройка заголовков
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	//инициализация графа
	var graph models.GraphNode

	//декодировать json
	err := json.NewDecoder(r.Body).Decode(&graph)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	//вызов функции вставки графа
	insertID := insertGraph(graph)

	//запись ответа
	res := response{
		ID:      insertID,
		Message: "Graph created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

//GetGraph вернет инфу о графе по его id
func GetGraph(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	//конвертация из строки в число
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	//вызов функции поиска по id
	graph, err := getGraph(int64(id))
	if err != nil {
		log.Fatalf("Unable to get graph. %v", err)
	}
	json.NewEncoder(w).Encode(graph)
}

//GetAllGraph вернет все графы
func GetAllGraph(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//поиск всех графов
	graphs, err := getAllGraph()
	if err != nil {
		log.Fatalf("Unable to get all graphs. %v", err)
	}
	json.NewEncoder(w).Encode(graphs)
}

//UpdateGraph обновление графа в таблице psql
func UpdateGraph(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var graph models.GraphNode
	err = json.NewDecoder(r.Body).Decode(&graph)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	//функция обновления графов
	updatedRows := updateGraph(int64(id), graph)
	//вывод сообщения об обновлении графов
	msg := fmt.Sprintf("Graph updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

//DeleteGraph удаление графа из таблицы
func DeleteGraph(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	//функция удаления графа
	deletedRows := deleteGraph(int64(id))
	msg := fmt.Sprintf("Graph deleted successfully. Total rows/record affected %v", deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

//------------------реализация функций

func insertGraph(graph models.GraphNode) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `insert into graphs(node,weight)values($1,$2)returning id`

	//сохраняем id
	var id int64

	//выполняем наш запрос
	err := db.QueryRow(sqlStatement, graph.Node).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)

	return id
}

func getGraph(id int64) (models.GraphNode, error) {
	db := createConnection()
	defer db.Close()
	var graph models.GraphNode
	sqlStatement := `select * from graphs where id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&graph.ID, &graph.Node, &graph.Weight)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return graph, nil
	case nil:
		return graph, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return graph, err
}

func getAllGraph() ([]models.GraphNode, error) {
	db := createConnection()
	defer db.Close()
	var graphs []models.GraphNode
	sqlStatement := `select * from graphs`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	//закрываем запрос
	defer rows.Close()

	//перебираем результаты
	for rows.Next() {
		var graph models.GraphNode
		err := rows.Scan(&graph.ID, &graph.Node, &graph.Weight)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		//добавляем граф в срез
		graphs = append(graphs, graph)
	}

	return graphs, err
}

func updateGraph(id int64, graph models.GraphNode) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `update graphs set node=$2 weight=$3 where id=$1`
	res, err := db.Exec(sqlStatement, id, graph.Node, graph.Weight)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	//проверка сколько строк затронуто
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func deleteGraph(id int64) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `delete from graphs where id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

//-----------реализация графа----------------

type Graph struct {
	VertexArray []*Vertex
}

type Vertex struct {
	Id      int64
	Visited bool
	AddEdge []*Edge
	Prev    *Vertex
}

type Edge struct {
	Source      *Vertex
	Destination *Vertex
	Weight      int64
}

func NewGraph() *Graph {
	return &Graph{
		make([]*Vertex, 0),
	}
}

func NewVertex(id int64) *Vertex {
	return &Vertex{
		Id:      id,
		Visited: false,
		AddEdge: make([]*Edge, 0),
	}
}

func NewEdge(source, destination *Vertex, weight int64) *Edge {
	return &Edge{
		source,
		destination,
		weight,
	}
}

func (G *Graph) AddVertexs(more ...*Vertex) {
	for _, vertex := range more {
		G.VertexArray = append(G.VertexArray, vertex)
	}
}

func (G *Graph) GetVertexByID(id int64) *Vertex {
	for _, vertex := range G.VertexArray {
		if vertex.Id == id {
			return vertex
		}
	}
	return nil
}

//Find the node with the id, or create it.
func (G *Graph) GetOrConst(id int64) *Vertex {
	vertex := G.GetVertexByID(id)
	if vertex == nil {
		vertex = NewVertex(id)
		G.AddVertexs(vertex)
	}
	return vertex
}

func (A *Vertex) AddEdges(more ...*Edge) {
	for _, edge := range more {
		A.AddEdge = append(A.AddEdge, edge)
	}
}

/*func ImportData(data int) *Graph {
	//	TODO
}*/

func (A *Vertex) GetAddEdg() chan *Edge {
	edgechan := make(chan *Edge)
	go func() {
		defer close(edgechan)
		for _, edge := range A.AddEdge {
			edgechan <- edge
		}
	}()

	return edgechan
}

func DFS(StartSource *Vertex) {
	if StartSource.Visited {
		return
	}

	StartSource.Visited = true
	fmt.Printf("%v ", StartSource.Id)

	for edge := range StartSource.GetAddEdg() {
		DFS(edge.Destination)
	}
}

const MAXWEIGHT = 1000000

type MinDistanceFromSource map[*Vertex]int64

func (G *Graph) Dijks(StartSource, TargetSource *Vertex) MinDistanceFromSource {
	D := make(MinDistanceFromSource)
	for _, vertex := range G.VertexArray {
		D[vertex] = MAXWEIGHT
	}
	D[StartSource] = 0

	for edge := range StartSource.GetAddEdg() {
		D[edge.Destination] = edge.Weight
	}
	CalculateDistance(StartSource, TargetSource, D)
	return D
}

func CalculateDistance(StartSource, TargetSource *Vertex, D MinDistanceFromSource) {
	for edge := range StartSource.GetAddEdg() {
		if D[edge.Destination] > D[edge.Source]+edge.Weight {
			D[edge.Destination] = D[edge.Source] + edge.Weight
			edge.Destination.Prev = edge.Source
		} else if D[edge.Destination] < D[edge.Source]+edge.Weight {
			continue
		}
		CalculateDistance(edge.Destination, TargetSource, D)
	}
}

func GraphPrint() {
	db := createConnection()
	defer db.Close()

	/*DFS(G1.ID)
		distmap:= G1.Dijks(G1,G2)

		fmt.Println()
	    for vertex,distance := range distmap {
			fmt.Println(vertex.Id, "===", distance)
			if vertex.Prev !=nil {
				fmt.Println(vertex.Id, " -> ", vertex.Prev.Id)
			}
		}*/

}
