package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/ashunasar/golang-students-crud-api/internal/config"
	"github.com/ashunasar/golang-students-crud-api/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	db *sql.DB
}

func New(cfg config.Config) (*Sqlite, error) {

	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{db: db}, nil

}

func (s Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.db.Prepare("INSERT INTO students(name,email,age) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil

}

func (s Sqlite) GetStudentById(id int64) (models.Student, error) {

	stmt, err := s.db.Prepare("Select * from students where id=? Limit 1")
	if err != nil {
		return models.Student{}, err
	}
	defer stmt.Close()

	var student models.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		return models.Student{}, fmt.Errorf("query error : %w", err)
	}

	return student, nil
}

func (s Sqlite) GetStudents() ([]models.Student, error) {
	var students []models.Student

	stmt, err := s.db.Prepare("Select * from students")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, fmt.Errorf("query error : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func (s Sqlite) UpdateStudent(student models.Student) (int64, error) {

	stmt, err := s.db.Prepare("UPDATE students set name = ?, email =?, age=? where id =?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(student.Name, student.Email, student.Age, student.Id)

	fmt.Println("results", result)

	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()

	fmt.Println("rowsAffected", rowsAffected)

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (s Sqlite) DeleteStudent(id int64) (int64, error) {

	stmt, err := s.db.Prepare("Delete From students where id=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)

	fmt.Println("results", &result)

	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()

	fmt.Println("rowsAffected", &rowsAffected)

	if err != nil {
		return 0, err
	}

	return id, nil
}
