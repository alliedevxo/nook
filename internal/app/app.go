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

// Notebooks
func (a *App) GetNotebooks() ([]db.Notebook, error) {
	return a.store.GetNotebooks()
}

func (a *App) InsertNotebook(title string) error {
	return a.store.InsertNotebook(title)
}

// Notes
func (a *App) GetNotes(notebookID int64) ([]db.Note, error) {
	return a.store.GetNotes(notebookID)
}

func (a *App) InsertNote(id int64, title, content string) error {
	return a.store.InsertNote(id, title, content)
}

func (a *App) MoveNote(noteID, notebookID int64) error {
	return a.store.MoveNote(noteID, notebookID)
}

func (a *App) UpdateNote(noteID int64, title, content string) error {
	return a.store.UpdateNote(noteID, title, content)
}
