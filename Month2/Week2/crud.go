package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// mini system to store details of altschool students

//type SlackName struct {
//	Full string
//	Display string
//}

type Student struct {
	ID     string
	Name   string
	Track  string
	Gender string
	Circle uint
	active bool
}

var studentsDB []Student

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentsDB)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, student := range studentsDB {
		if student.ID == params["id"] {
			json.NewEncoder(w).Encode(student)
			return
		}
	}
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, student := range studentsDB {
		if student.ID == params["id"] {
			fmt.Println()
			studentsDB = append(studentsDB[:index], studentsDB[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(studentsDB)
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	student.ID = strconv.Itoa(rand.Intn(1000000000))
	studentsDB = append(studentsDB, student)
	json.NewEncoder(w).Encode(student)
}

func appHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I'm alive 🍾")
	return
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", appHealth).Methods("GET")
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students", createStudent).Methods("POST")
	r.HandleFunc("/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	fmt.Printf("Starting altschool server at port 5000 \n")
	log.Fatal(http.ListenAndServe(":5000", r))
}
