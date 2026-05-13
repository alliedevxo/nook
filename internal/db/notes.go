package db

import (
	"fmt"
)

func (store *Store) InsertNote(notebookId int64, title, body string) error {
	const insertNoteSql = `
	INSERT INTO notes (notebook_id, title, body)
	VALUES (?, ?, ?)
	`

	if _, err := store.db.Exec(insertNoteSql, notebookId, title, body); err != nil {
		return err
	}

	return nil
}

func (store *Store) GetNotes() error {
	const viewNotesTableSql = `
	SELECT * FROM notes;
	`

	rows, err := store.db.Query(viewNotesTableSql)
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