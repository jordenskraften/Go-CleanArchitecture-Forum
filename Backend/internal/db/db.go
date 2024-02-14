package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/jordenskraften/Go-CleanArchitecture-Forum/internal/config"
)

type DataBase struct {
	Cfg    *config.Config
	Logger *slog.Logger
	DB     *sql.DB
}

func NewDB(logger *slog.Logger, cfg *config.Config) *DataBase {
	dbConfig := &cfg.DB
	connstring := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		dbConfig.User, dbConfig.DbName, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.SSLmode)

	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	}

	DataBase := &DataBase{
		Cfg:    cfg,
		Logger: logger,
		DB:     db,
	}
	return DataBase
}
