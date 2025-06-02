package global

import (
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

var (
	Logger         *log.Logger
	DB             *pgxpool.Pool
	WhileListPaths map[string]string
)
