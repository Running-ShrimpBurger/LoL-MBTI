package model

import (
	"database/sql"
	"fmt"

	"github.com/Lol-MBTI/secret"
)

type mysqlHandler struct {
	db *sql.DB
}

func (m *mysqlHandler) Close() {
	m.db.Close()
}

func newMysqlHandler() DBHandler {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", secret.GetDBHost(), secret.GetDBPort(), secret.GetDBUser(), secret.GetDBPassword(), "mysql")

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS champions (
			id INTEGER NOT NULL,
			name TEXT,
			mbti TEXT,
			PRIMARY KEY (id)
		);`)
	statement.Exec()
	return &mysqlHandler{database}
}
