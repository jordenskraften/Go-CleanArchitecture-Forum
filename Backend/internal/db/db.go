package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/jordenskraften/Go-CleanArchitecture-Forum/internal/config"
	_ "github.com/lib/pq"
)

type DataBase struct {
	Cfg    *config.Config
	Logger *slog.Logger
	DB     *sql.DB
}

func NewDB(logger *slog.Logger, cfg *config.Config) *DataBase {
	dbConfig := &cfg.DB
	connstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)

	db, err := sql.Open("postgres", connstring)
	if err != nil {
		logger.Info("Connecting to database",
			slog.Any("db error: ", err),
		)
	}

	DataBase := &DataBase{
		Cfg:    cfg,
		Logger: logger,
		DB:     db,
	}

	err = DataBase.DB.Ping()
	if err != nil {
		panic(err)
	}
	logger.Info("DB Status",
		slog.String("succesfully conected to db=", cfg.DB.Dbname),
	)

	return DataBase
}
