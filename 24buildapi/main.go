package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for courses - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// middleware, helper - file
func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

}

// Controllers - file

// Serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page")
	w.Write([]byte("<h1>Welcome to the home page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// Get the course id from the URL
	params := mux.Vars(r)
	fmt.Println("parmas - ", params)

	// Loop through the courses to find the course with the id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// If the course is not found
	json.NewEncoder(w).Encode("Course not found")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		// return
	}

	// What about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	// Check if the course is empty
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate a unique id, string
	// append course to courses
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	parmas := mux.Vars(r)

	// loop, id, remove, add with my Id
	for index, course := range courses {
		if course.CourseId == parmas["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = parmas["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// TODO: handle when id is not found 
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Loop, id, remove

	for index, course := range courses{
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted")
			return
		}
	}

}