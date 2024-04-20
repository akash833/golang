package main

type course struct {
	CourseId    string `json:"courseId"`
	CourseName  string `json:"coursename"`
	CoursePrice int64  `json:"coursePrice"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// fake DB
var Courses []course

// middleware & helper - file
func (c *course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}
