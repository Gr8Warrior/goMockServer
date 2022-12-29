package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

	//return c.CouserId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func main() {
	fmt.Println("Mock server in json")
	r := mux.NewRouter()

	//seeding of data
	courses = append(courses, Course{"1", "Go Lang", 299, &Author{"Ken", "go.dev"}})
	courses = append(courses, Course{"2", "Java Basics", 199, &Author{"Shailu", "gr8warrior.com"}})

	//routing
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/course", addCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourseById).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourseById).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
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

	//grab id from request
	params := mux.Vars(r)
	fmt.Printf("type : %T, Value = %v\n", params, params)

	//loop through the couses and fund the matching id and return the response
	for _, course := range courses {
		if course.CouserId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")

}

func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add a course")
	w.Header().Set("Content-type", "application/json")

	//what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	//what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return

	}

	//generate unique id, string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CouserId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

}

func updateCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")
	w.Header().Set("Content-type", "application/json")

	//first - grab id from request
	params := mux.Vars(r)

	//loop, id, remove, add with my id(from request)
	for index, course := range courses {
		//check id
		if course.CouserId == params["id"] {
			//remove
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CouserId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}

	// send a response when id is not find
	json.NewEncoder(w).Encode("courseid not found")

}

func deleteCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course by id")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove
	for index, course := range courses {
		//check id
		if course.CouserId == params["id"] {
			//remove
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted!!!")
			break
		}
	}

	json.NewEncoder(w).Encode("CourseId not found")
}
