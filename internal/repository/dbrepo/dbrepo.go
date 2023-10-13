package dbrepo

import (
	"bookings/internal/config"
	"bookings/internal/repository"
	"database/sql"
)

type mysqlDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewMysqlRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &mysqlDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingsRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}
