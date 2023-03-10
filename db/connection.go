package db

import (
	"database/sql"
	"fmt"

	"github.com/downtownit/golang/api-pgsql/configs"
	_ "github.com/lib/pq"
)

func OpemConnection() (*sql.DB, error) {
	fmt.Println("OpemConnectionOpemConnectionOpemConnection", configs.GetDB())
	conf := configs.GetDB()
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	return conn, err
}
