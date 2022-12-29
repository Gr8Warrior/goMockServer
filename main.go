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
