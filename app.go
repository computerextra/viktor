package main

import (
	"context"
	"fmt"
	"log"
	"viktor/ent"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
)

// App struct
type App struct {
	ctx         context.Context
	environment *Env
	db          *ent.Client
}

// NewApp creates a new App application struct
func NewApp(env *Env) *App {

	client, err := ent.Open(dialect.MySQL, "d043fb41:9Vm7chP99iD4uSwupQKs@tcp(computer-extra.de:3306)/d043fb41?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
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
