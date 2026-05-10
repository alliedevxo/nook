package db

import (
	"database/sql"
	"fmt"
	"nook/internal/filesystem"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {
	path, err := filesystem.GetUserConfigDir()
	if err != nil {
		return nil, err
	}

	dsn := path + "?_foreign_keys=on"

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	if err := initializeTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func initializeTables(db *sql.DB) error {
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
		body TEXT,
		FOREIGN KEY (notebook_id) REFERENCES notebooks(id) ON DELETE CASCADE
	);
	`
	
	tx, err := db.Begin()

	if err != nil {
		dbErr := fmt.Errorf("start transaction: %w", err)
		return dbErr
	}
	defer tx.Rollback()

	if _, err := tx.Exec(createNotebookSql); err != nil {
		dbErr := fmt.Errorf("create notebooks table: %w", err)
		return dbErr
	}

	if _, err := tx.Exec(createNotesSql); err != nil {
		dbErr := fmt.Errorf("create notes table: %w", err)
		return dbErr
	}

	return tx.Commit()
}

func InsertNotebook(db *sql.DB, title string) error {
	const viewNotebookTableSql = `
	INSERT INTO notebooks (title)
	VALUES (?)
	`

	if _, err := db.Exec(viewNotebookTableSql, title); err != nil {
		return err
	}

	return nil
}

func ViewNotebooks(db *sql.DB) error {
	const viewNotebookTableSql = `
	SELECT * FROM notebooks;
	`

	rows, err := db.Query(viewNotebookTableSql)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var title string
		if err := rows.Scan(&id, &title); err != nil {
			return err
		}

		fmt.Printf("Id: %d; Title: %s\n", id, title)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	
	return nil
}

func InsertNote(db *sql.DB, notebook_id int64, title, body string) error {
	const viewNotebookTableSql = `
	INSERT INTO notes (notebook_id, title, body)
	VALUES (?, ?, ?)
	`

	if _, err := db.Exec(viewNotebookTableSql, notebook_id, title, body); err != nil {
		return err
	}

	return nil
}

func ViewNotes(db *sql.DB) error {
	const viewNotesTableSql = `
	SELECT * FROM notes;
	`

	rows, err := db.Query(viewNotesTableSql)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var notebook_id int64
		var title string
		var body string
		if err := rows.Scan(&id, &notebook_id, &title, &body); err != nil {
			return err
		}

		fmt.Printf("Id: %d; Notebook Id: %d; Title: %s; Body: %s\n", id, notebook_id, title, body)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	
	return nil
}


