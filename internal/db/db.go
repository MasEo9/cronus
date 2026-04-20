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

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//	 TABLE CREATION
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (db *CronusDB) SessionTableCreate() error {
	createSessionTableQuery := `CREATE TABLE IF NOT EXISTS sessions (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			project_id INTEGER NOT NULL, 
			date TEXT NOT NULL, 
			time_start DATETIME NOT NULL, 
			time_end DATETIME, 
			elapsed_time REAL, 
			hours REAL
	);`
	_, err := db.db.Exec(createSessionTableQuery)
	return err
}

func (db *CronusDB) ProjectTableCreate() error {
	createProjectTableQuery := `CREATE TABLE IF NOT EXISTS projects (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			project_name TEXT NOT NULL 
	);`
	_, err := db.db.Exec(createProjectTableQuery)
	return err
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//	PROJECT QUERIES
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (db *CronusDB) InsertProject(name string) error {

	_, err := db.db.Exec(`INSERT INTO projects(project_name) VALUES(?);`, name)

	return err
}

func (db *CronusDB) ListProjects() ([]models.Project, error) {
	var projects []models.Project

	// scan binds the output values back to the model through a pointer
	query := `SELECT id, project_name FROM projects ORDER BY id ASC;`
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := &models.Project{}
		rows.Scan(&s.ID, &s.ProjectName)
		projects = append(projects, *s)
	}

	return projects, err
}

func (db *CronusDB) SearchProjects(projectName string) (models.Project, error) {
	var projects models.Project

	query := `SELECT id, project_name FROM projects WHERE project_name = ?;`
	p := db.db.QueryRow(query, projectName)
	p.Scan(&projects.ID, &projects.ProjectName)

	return projects, nil
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//		SESSION QUERIES
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (db *CronusDB) GetActiveSession() ([]models.Session, error) {
	// data is saved in this var
	var sessions []models.Session

	// scan binds the output values back to the model through a pointer
	query := `SELECT s.id, s.time_start, p.id as project_id, p.project_name FROM sessions s
			LEFT JOIN projects p on p.id = s.project_id 
		WHERE s.time_end IS NULL ORDER BY s.id DESC;`
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := &models.Session{}
		rows.Scan(&s.ID, &s.TimeStart, &s.ProjectID, &s.ProjectName)
		sessions = append(sessions, *s)
	}

	return sessions, err
}

func (db *CronusDB) InsertSession(project_id int, date string, timeStart time.Time) error {

	_, err := db.db.Exec(`INSERT INTO sessions(project_id, date, time_start) VALUES(?,?,?);`, project_id, date, timeStart)

	return err
}

func (db *CronusDB) StopSession(timeStop time.Time, timeElapsed time.Duration, hours float32, id int) error {
	_, err := db.db.Exec("UPDATE sessions SET time_end = ?, elapsed_time = ?, hours = ? WHERE id = ?;", timeStop, timeElapsed, hours, id)

	return err
}
