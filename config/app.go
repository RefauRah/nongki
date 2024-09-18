package config

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	Db     *sqlx.DB
	Logger *logrus.Logger
}
