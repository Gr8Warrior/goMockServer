package main

import "fmt"

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
