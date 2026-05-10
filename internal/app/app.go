package app

import (
	"context"
	"database/sql"
	"fmt"
)

type App struct {
	ctx context.Context
	db *sql.DB
}

func New(db *sql.DB) *App {
	return &App{db: db}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello from nook, %s!", name)
}
