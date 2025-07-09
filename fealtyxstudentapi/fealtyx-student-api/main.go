package main

import (
	"net/http"
	"sync"

	"fealtyx-student-api/models"

	"github.com/gorilla/mux"
)

var (
	students = make(map[int]models.Student)
	mu       sync.Mutex
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students", CreateStudent).Methods("POST")
	r.HandleFunc("/students", GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{id}", GetStudent).Methods("GET")
	r.HandleFunc("/students/{id}", UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students/{id}/summary", GenerateSummary).Methods("GET")

	// Add this line for confirmation:
	println("✅ Server is running at http://localhost:8080")

	http.ListenAndServe(":8080", r)
}
