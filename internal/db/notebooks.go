package db

type Notebook struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func (store *Store) InsertNotebook(title string) error {
	const insertNotebookSql = `
	INSERT INTO notebooks (title)
	VALUES (?)
	`

	if _, err := store.db.Exec(insertNotebookSql, title); err != nil {
		return err
	}

	return nil
}

func (store *Store) GetNotebooks() ([]Notebook, error) {
	const getNotebookTableSql = `
	SELECT * FROM notebooks;
	`

	rows, err := store.db.Query(getNotebookTableSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Notebook

	for rows.Next() {
		var id int64
		var title string
		if err := rows.Scan(&id, &title); err != nil {
			return nil, err
		}

		out = append(out, Notebook{ID: id, Title: title})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return out, nil
}