// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Assignment struct {
	ID               int64  `json:"id"`
	ModuleID         int64  `json:"module_id"`
	Description      string `json:"description"`
	Content          []byte `json:"content"`
	AssignmentTypeID int64  `json:"assignment_type_id"`
}

type AssignmentsType struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
}

type Course struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CourseCategory struct {
	CourseID   pgtype.Int8 `json:"course_id"`
	CategoryID pgtype.Int8 `json:"category_id"`
}

type Deadline struct {
	ID            int64       `json:"id"`
	AssignmentsID int64       `json:"assignments_id"`
	Deadline      interface{} `json:"deadline"`
	UserID        int64       `json:"user_id"`
}

type Enrollment struct {
	ID         int64       `json:"id"`
	EnrolledOn pgtype.Date `json:"enrolled_on"`
	CourseID   int64       `json:"course_id"`
	UserID     int64       `json:"user_id"`
}

type Module struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	CourseID int64  `json:"course_id"`
}

type Rating struct {
	ID        int64       `json:"id"`
	CommentID pgtype.Int8 `json:"comment_id"`
	Rate      bool        `json:"rate"`
}

type Submission struct {
	ID            int64  `json:"id"`
	AssignmentsID int64  `json:"assignments_id"`
	DeadlineID    int64  `json:"deadline_id"`
	Delay         int32  `json:"delay"`
	Content       []byte `json:"content"`
}

type Thread struct {
	ID       int64       `json:"id"`
	ModuleID int64       `json:"module_id"`
	Title    string      `json:"title"`
	Content  pgtype.Text `json:"content"`
	UserID   int64       `json:"user_id"`
}

type User struct {
	ID         int64       `json:"id"`
	Login      string      `json:"login"`
	Password   string      `json:"password"`
	Surname    pgtype.Text `json:"surname"`
	Firstname  pgtype.Text `json:"firstname"`
	UserRoleID int32       `json:"user_role_id"`
}

type UserRole struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}
