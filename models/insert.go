package models

import (
	"github.com/downtownit/golang/api-pgsql/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpemConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3)  RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(id)
	return
}
