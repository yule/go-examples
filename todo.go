package main


import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todos", CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	// Extract the todo ID from the request parameters
	params := mux.Vars(r)
	for _, item := range todos {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			var updatedTodo Todo
			_ = json.NewDecoder(r.Body).Decode(&updatedTodo)
			todos[index] = updatedTodo
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTodo)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Todo deleted")
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}
