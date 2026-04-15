package db

import (
	"database/sql"
	"fmt"
	"time"

	"cronus/internal/models"
	_ "modernc.org/sqlite"
)

type CronusDB struct {
	db *sql.DB
}

func New(dbPath string) (*CronusDB, error) {
	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &CronusDB{db: conn}, nil
}

func (db *CronusDB) SessionTableCreate() error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS sessions (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			project_name TEXT NOT NULL, 
			date TEXT NOT NULL, 
			time_start DATETIME NOT NULL, 
			time_end DATETIME, 
			elapsed_time REAL, 
			hours REAL
	);`
	_, err := db.db.Exec(createTableQuery)
	return err
}

func (db *CronusDB) InsertSession(name string, date string, timeStart time.Time) error {

	_, err := db.db.Exec(`INSERT INTO sessions(project_name, date, time_start) VALUES(?,?,?);`, name, date, timeStart)

	return err
}

func (db *CronusDB) GetActiveSession(project_name string) (models.Session, error) {
	// data is saved in this var
	var s models.Session

	// scan binds the output values back to the model through a pointer
	err := db.db.QueryRow(`SELECT id, time_start FROM sessions 
		WHERE project_name = ? AND time_end IS NULL ORDER BY id DESC;`, project_name).Scan(&s.ID, &s.TimeStart)

	fmt.Println("Active Session", s)
	return s, err
}

func (db *CronusDB) StopSession(timeStop time.Time, timeElapsed time.Duration, hours float32, id int) error {
	_, err := db.db.Exec("UPDATE sessions SET time_end = ?, elapsed_time = ?, hours = ? WHERE id = ?;", timeStop, timeElapsed, hours, id)

	return err
}
