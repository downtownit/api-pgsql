package models

import "github.com/downtownit/golang/api-pgsql/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpemConnection()
	if err != nil {
		return 0, err
	}
	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
