package internal

import (
	"database/sql"
	mapper "github.com/birkirb/loggers-mapper-logrus"
	"gopkg.in/birkirb/loggers.v1"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"os"
)

type Resolver struct {
	config *Config

	logger loggers.Contextual
	db *sql.DB
}

func NewResolver(c *Config) *Resolver {
	return &Resolver{
		config: c,
	}
}

type LoggerResolver interface {
	ResolveLogger() loggers.Contextual
}

func (r *Resolver) ResolveLogger() loggers.Contextual {
	if r.logger == nil {
		l := logrus.New()
		l.Out = os.Stdout
		l.Level = logrus.InfoLevel
		l.SetFormatter(&logrus.JSONFormatter{})
		r.logger = mapper.NewLogger(l)
	}
	return r.logger
}

func (r *Resolver) ResolveDB() *sql.DB {
	if r.db == nil {
		db, err := sql.Open("mysql", r.config.DB.GetConnectionString())
		if err != nil {
			panic(err)
		}

		// find a place for these
		db.SetMaxIdleConns(5)
		db.SetMaxOpenConns(5)

		return db
	}
	return r.db
}
