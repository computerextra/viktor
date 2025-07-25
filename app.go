package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// App struct
type App struct {
	ctx         context.Context
	environment *Env
	db          *sql.DB
}

// NewApp creates a new App application struct
func NewApp(env *Env) *App {
	client, err := sql.Open("mysql", env.DATABASE_URL)
	if err != nil {
		panic(err)
	}

	return &App{
		environment: env,
		db:          client,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
