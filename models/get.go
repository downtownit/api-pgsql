package models

import "github.com/downtownit/golang/api-pgsql/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpemConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM todos WHERE id = $1`, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
