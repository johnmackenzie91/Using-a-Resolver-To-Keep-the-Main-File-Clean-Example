package main

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/birkirb/loggers.v1"
	"os"
	mapper "github.com/birkirb/loggers-mapper-logrus"
)

func main()  {
	dbConf := GetDBConfig()
	db, err := sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConf.dbUser, dbConf.dbPass, dbConf.dbHost, dbConf.dbPort, dbConf.dbName) )
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)

	// init the repo with the help of the resolver
	repo := Repo{
		DB: db,
	}

	// init the app
	a := App{
		Logger: GetLogger(),
		Repo: repo,
	}
	a.Run()
}

type Config struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
}


func GetDBConfig() Config {
	return Config{
		dbUser: getMandidatoryEnvVar("DB_USER"),
		dbPass: getMandidatoryEnvVar("DB_PASS"),
		dbHost: getMandidatoryEnvVar("DB_HOST"),
		dbPort: getMandidatoryEnvVar("DB_PORT"),
		dbName: getMandidatoryEnvVar("DB_NAME"),
	}
}

func getMandidatoryEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("mandidatory env var not set %s", key))
	}
	return val
}


func GetLogger () loggers.Contextual {
	l := logrus.New()
	l.Out = os.Stdout
	l.Level = logrus.InfoLevel
	l.SetFormatter(&logrus.JSONFormatter{})
	return mapper.NewLogger(l)
}
