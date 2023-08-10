package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlCourseRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func (r *MysqlCourseRepository) Save(ctx context.Context, course domain.Course) error {
	stmt, err := r.db.Prepare("INSERT INTO courses(id, name) VALUES(?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err = stmt.ExecContext(ctxTimeout, course.Id(), course.Name())
	if err != nil {
		return err
	}

	return nil
}

func NewMysqlCourseRepository(user, password, host, port, database string, timeout time.Duration) (*MysqlCourseRepository, error) {
	connectionString := gerConnectionString(user, password, host, port, database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &MysqlCourseRepository{
		db:        db,
		dbTimeout: timeout,
	}, nil
}

func gerConnectionString(user, password, host, port, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
}
