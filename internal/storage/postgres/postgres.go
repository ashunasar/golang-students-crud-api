package postgres

import (
	"database/sql"
	"fmt"

	"github.com/ashunasar/golang-students-crud-api/internal/config"
	"github.com/ashunasar/golang-students-crud-api/internal/models"
	_ "github.com/lib/pq" // Importing the Postgres driver
)

type Postgres struct {
	db *sql.DB
}

func New(cfg config.Config) (*Postgres, error) {
	// Assuming `cfg.StoragePath` is the connection string for PostgreSQL
	db, err := sql.Open("postgres", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	// Create the table if not exists (adjust for PostgreSQL)
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT,
			age INTEGER
		)
	`)

	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (p Postgres) CreateStudent(name string, email string, age int) (int64, error) {
	// Use INSERT INTO with parameterized query
	stmt, err := p.db.Prepare("INSERT INTO students(name, email, age) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var lastId int64
	err = stmt.QueryRow(name, email, age).Scan(&lastId)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (p Postgres) GetStudentById(id int64) (models.Student, error) {
	stmt, err := p.db.Prepare("SELECT * FROM students WHERE id=$1 LIMIT 1")
	if err != nil {
		return models.Student{}, err
	}
	defer stmt.Close()

	var student models.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		return models.Student{}, fmt.Errorf("query error: %w", err)
	}

	return student, nil
}

func (p Postgres) GetStudents() ([]models.Student, error) {
	var students []models.Student

	stmt, err := p.db.Prepare("SELECT * FROM students")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
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

func (p Postgres) UpdateStudent(student models.Student) (int64, error) {
	stmt, err := p.db.Prepare("UPDATE students SET name=$1, email=$2, age=$3 WHERE id=$4")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(student.Name, student.Email, student.Age, student.Id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (p Postgres) DeleteStudent(id int64) (int64, error) {
	stmt, err := p.db.Prepare("DELETE FROM students WHERE id=$1")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
