package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlHandler struct {
	db *sql.DB
}

func (m *mysqlHandler) Close() {
	m.db.Close()
}

func newMysqlHandler() DBHandler {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/mbti-db", secret.GetDBUser(), secret.GetDBPassword(), secret.GetDBHost(), secret.GetDBPort())

	// database, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	panic(err)
	// }
	// defer database.Close()

	// err = database.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// statisticsStatement, _ := database.Prepare(
	// 	`CREATE TABLE IF NOT EXISTS statistics (
	// 		id INT NOT NULL AUTO_INCREMENT,
	// 		line VARCHAR(30),
	// 		mbti VARCHAR(4),
	// 		CONSTRAINT statistics_PK PRIMARY KEY(id)
	// 	);`)
	// statisticsStatement.Exec()
	return &mysqlHandler{}
}
