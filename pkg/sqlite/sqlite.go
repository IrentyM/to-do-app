package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func NewSQLiteConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./todo.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS tasks (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            done BOOLEAN NOT NULL DEFAULT 0,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            deadline TIMESTAMP,
            priority TEXT CHECK(priority IN ('low', 'medium', 'high')) DEFAULT 'low'
        );
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}
