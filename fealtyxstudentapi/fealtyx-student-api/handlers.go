package main

import (
	"encoding/json"
	"fealtyx-student-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	json.NewDecoder(r.Body).Decode(&student)

	if student.ID == 0 || student.Name == "" || student.Age <= 0 || student.Email == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	if _, exists := students[student.ID]; exists {
		http.Error(w, "Student ID already exists", http.StatusBadRequest)
		return
	}

	students[student.ID] = student
	json.NewEncoder(w).Encode(student)
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var all []models.Student
	for _, s := range students {
		all = append(all, s)
	}
	json.NewEncoder(w).Encode(all)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	student, exists := students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	var updated models.Student
	json.NewDecoder(r.Body).Decode(&updated)

	if updated.Name == "" || updated.Age <= 0 || updated.Email == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	_, exists := students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	updated.ID = id
	students[id] = updated
	json.NewEncoder(w).Encode(updated)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	mu.Lock()
	defer mu.Unlock()

	_, exists := students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	delete(students, id)
	w.WriteHeader(http.StatusNoContent)
}
