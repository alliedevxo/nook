package app

import (
	"context"

	"nook/internal/db"
)

type App struct {
	ctx context.Context
	store *db.Store
}

func New(store *db.Store) *App {
	return &App{store: store}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetNotebooks() ( []db.Notebook, error) {
	return a.store.GetNotebooks()
}
