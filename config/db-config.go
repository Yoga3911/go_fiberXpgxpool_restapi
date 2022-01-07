package config

import "database/sql"

func DBConnection() *sql.DB {
	sql, err := sql.Open("postgres", "dbname=api user=yoga password=yyooggaa2020 port=5432 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	return sql
}