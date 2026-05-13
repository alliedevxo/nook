package db

import "database/sql"

type Note struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (store *Store) InsertNote(notebookID int64, title, content string) error {
	const insertNoteSql = `
	INSERT INTO notes (notebook_id, title, content)
	VALUES (?, ?, ?)
	`

	if _, err := store.db.Exec(insertNoteSql, notebookID, title, content); err != nil {
		return err
	}

	return nil
}

func (store *Store) UpdateNote(noteID int64, title, content string) error {
	const updateNoteSql = `
	UPDATE notes
	SET title = ?, content = ?
	WHERE id = ?
	`

	res, err := store.db.Exec(updateNoteSql, title, content, noteID)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (store *Store) MoveNote(noteID, notebookID int64) error {
	const moveNoteSql = `
	UPDATE notes
	SET notebook_id = ?
	WHERE id = ?
	`

	res, err := store.db.Exec(moveNoteSql, notebookID, noteID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (store *Store) GetNotes(notebookID int64) ([]Note, error) {
	const getNotesTableSql = `
	SELECT id, title, content FROM notes
	WHERE notebook_id = ?;
	`

	rows, err := store.db.Query(getNotesTableSql, notebookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Note

	for rows.Next() {
		var id int64
		var title string
		var content string
		if err := rows.Scan(&id, &title, &content); err != nil {
			return nil, err
		}

		out = append(out, Note{ID: id, Title: title, Content: content})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return out, nil
}