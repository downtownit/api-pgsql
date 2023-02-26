package models

import "github.com/downtownit/golang/api-pgsql/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpemConnection()
	if err != nil {
		return 0, err
	}
	res, err := conn.Exec(`UPDATE todos SET title=$1, description=$2, done=$3 WHERE id=$4`, todo.Title, todo.Description, todo.Done, todo.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
