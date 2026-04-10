package storage

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

type Command struct {
	Command   string
	Directory string
	ExitCode  int
	StartTime int64
	EndTime   int64
}

func InitDB() (*sql.DB, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	path := home + "/.cmdock.db"
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := createTable(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS commands(
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		command    TEXT,
		directory  TEXT,
		exit_code  INTEGER,
		start_time INTEGER,
		end_time   INTEGER
	);`
	_, err := db.Exec(query)
	return err
}

func InsertCommand(db *sql.DB, c Command) error {
	query := `
	INSERT INTO commands(command, directory, exit_code, start_time, end_time)
	VALUES(?, ?, ?, ?, ?)`
	_, err := db.Exec(query, c.Command, c.Directory, c.ExitCode, c.StartTime, c.EndTime)
	return err
}

func ShowCommands(db *sql.DB) ([]Command, error) {
	query := `
	SELECT command, directory, exit_code, start_time, end_time
	FROM commands
	ORDER BY start_time DESC
	LIMIT 20`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Command
	for rows.Next() {
		var c Command
		err := rows.Scan(&c.Command, &c.Directory, &c.ExitCode, &c.StartTime, &c.EndTime)
		if err != nil {
			return nil, err
		}
		results = append(results, c)
	}

	return results, rows.Err()
}