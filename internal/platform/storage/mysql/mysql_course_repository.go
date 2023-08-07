package mysql

import (
	"database/sql"
	"fmt"
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlCourseRepository struct {
	db *sql.DB
}

func (r *MysqlCourseRepository) Save(course domain.Course) error {
	stmt, err := r.db.Prepare("INSERT INTO course(id, name) VALUES(?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Id(), course.Name())
	if err != nil {
		return err
	}

	return nil
}

func NewMysqlCourseRepository(user, password, host, port, database string) (*MysqlCourseRepository, error) {
	connectionString := gerConnectionString(user, password, host, port, database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &MysqlCourseRepository{
		db: db,
	}, nil
}

func gerConnectionString(user, password, host, port, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
}
