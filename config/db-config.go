package config

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func DBConnection() *pgxpool.Pool {
	sql, err := pgxpool.Connect(context.Background(), "postgres://yoga:yyooggaa2020@localhost:5432/api")
	if err != nil {
		panic(err.Error())
	}
	return sql
}