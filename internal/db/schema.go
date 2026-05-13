package db

import (
	"fmt"
)

func (store *Store) InitializeTables() error {
	const createNotebookSql = `
	CREATE TABLE IF NOT EXISTS notebooks(
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL
	);
	`

	const createNotesSql = `
	CREATE TABLE IF NOT EXISTS notes(
		id INTEGER PRIMARY KEY,
		notebook_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		content TEXT,
		FOREIGN KEY (notebook_id) REFERENCES notebooks(id) ON DELETE CASCADE
	);
	`

	tx, err := store.db.Begin()

	if err != nil {
		return fmt.Errorf("start transaction: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.Exec(createNotebookSql); err != nil {
		return fmt.Errorf("create notebooks table: %w", err)
	}

	if _, err := tx.Exec(createNotesSql); err != nil {
		return fmt.Errorf("create notes table: %w", err)
	}

	return tx.Commit()
}