package storage

import "github.com/ashunasar/golang-students-crud-api/internal/models"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)

	GetStudentById(id int64) (models.Student, error)

	GetStudents() ([]models.Student, error)

	UpdateStudent(student models.Student) (int64, error)

	DeleteStudent(id int64) (int64, error)
}
