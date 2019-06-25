package main

import (
	"database/sql"
	"gopkg.in/birkirb/loggers.v1"
	"no_vcs/me/resolver/internal"
)

func main()  {
	// init the environment configuration
	c := internal.NewConfiguration()
	// init the resolver
	r := internal.NewResolver(c)


	// init the repo with the help of the resolver
	repo := Repo{
		DB: r.ResolveDB(),
	}

	// init the app
	a := App{
		Logger: r.ResolveLogger(),
		Repo: repo,
	}
	a.Run()
}

type Repo struct {
	DB *sql.DB
}

func (r Repo) CheckConnection() error {
	return r.DB.Ping()
}

type App struct {
	Logger loggers.Contextual
	Repo Repo
}

func (a App) Run () {
	a.Logger.Info("running something")
	// ...
	if err := a.Repo.CheckConnection(); err != nil {
		panic(err)
	}
	//
	a.Logger.Info("operation complete")
}
