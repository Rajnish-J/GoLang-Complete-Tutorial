package main

import (
	"fmt"
	"strings"
)

// Sample student-course mapping
var studentCourses = map[string]string{
	"Alice":   "Mathematics",
	"Bob":     "Physics",
	"Charlie": "Chemistry",
	"David":   "physics",
	"Eve":     "Mathematics",
}

// API #1: Print all students with their enrolled courses
func PrintAllStudentsAndCourses() {
	fmt.Println("All Students and Their Courses:")
	for student, course := range studentCourses {
		fmt.Printf("%s - %s\n", student, course)
	}
}

// API #2: Print students by course (case-insensitive)
func FindStudentsByCourse(course string) {
	course = strings.ToLower(course)
	fmt.Printf("\nStudents enrolled in %s:\n", course)
	found := false
	for student, enrolledCourse := range studentCourses {
		if strings.ToLower(enrolledCourse) == course {
			fmt.Println(student)
			found = true
		}
	}
	if !found {
		fmt.Println("No students found for this course.")
	}
}

func main() {
	// Test API 1
	PrintAllStudentsAndCourses()

	// Test API 2
	FindStudentsByCourse("Physics")
	FindStudentsByCourse("Mathematics")
	FindStudentsByCourse("Biology")
}
