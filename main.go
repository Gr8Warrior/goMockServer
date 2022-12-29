package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model for course -file
type Course struct {
	CouserId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper -file
func (c *Course) isEmpty() bool {

	return c.CouserId == "" && c.CourseName == ""

}

func main() {
	fmt.Println("Mock server in json")
}

//controllers - file

// server home route
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to Mock API Server </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course by id")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses[0])

	//grab id from request
	params := mux.Vars(r)
	fmt.Printf("type : %T, Value = %v\n", params, params)

	//loop through the couses and fund the matching id and return the response
	for _, course := range courses {
		if course.CouserId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
		json.NewEncoder(w).Encode("No Course found with given id %d", params["id"])
		return
	}

}
