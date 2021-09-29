package model

import (
	"database/sql"
	"fmt"

	"github.com/Lol-MBTI/secret"
	_ "github.com/go-sql-driver/mysql"
)

type mysqlHandler struct {
	db *sql.DB
}

func (m *mysqlHandler) Close() {
	m.db.Close()
}

func newMysqlHandler() DBHandler {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/", secret.GetDBUser(), secret.GetDBPassword(), "tcp", secret.GetDBHost(), secret.GetDBPort())

	database, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	_, err = database.Exec("CREATE DATABASE test")
	if err != nil {
		panic(err)
	}

	_, err = database.Exec("USE test")
	if err != nil {
		panic(err)
	}

	_, err = database.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}

	champStatement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS champions (
			id INT NOT NULL AUTO_INGREMENT,
			name VARCHAR(30),
			mbti VARCHAR(4),
			CONSTRAINT champions_PK PRIMARY KEY(id)
		);`)
	compatibilityStatement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS compatibility (
			id INT NOT NULL AUTO_INGREMENT,
			mbti VARCHAR(4),
			good VARCHAR(4),
			bad VARCHAR(4),
			CONSTRAINT champions_PK PRIMARY KEY(id)
		);`)

	champStatement.Exec()
	compatibilityStatement.Exec()
	return &mysqlHandler{database}
}
