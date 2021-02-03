package internal

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type LogRecord struct {
	Id int
	Message string
	Time time.Time
}

type DbManager struct {
	database *sql.DB
}

func NewDbManager(file string) (error, *DbManager) {
	database, err := sql.Open("sqlite3", file)
	if err != nil {
		return err, nil
	}

	stmt, err := database.Prepare("CREATE TABLE IF NOT EXISTS logs (id INTEGER PRIMARY KEY, message TEXT NOT NULL, time DATETIME)")
	if err != nil {
		return err, nil
	}

	_, err = stmt.Exec()
	if err != nil {
		return err, nil
	}

	return nil, &DbManager {
		database: database,
	}
}

func (dbManager *DbManager) Insert(message string, time time.Time) (sql.Result, error) {
	stmt, err := dbManager.database.Prepare("INSERT INTO logs (message, time) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	return stmt.Exec(message, time)
}

func (dbManager *DbManager) List(date time.Time) (*sql.Rows, error) {
	return dbManager.database.Query("SELECT id, message, time FROM logs WHERE DATE(time) = DATE(?)", date)
}

func (dbManager *DbManager) GetLastInsertDate() (time.Time, error) {
	// subselect coz column affinity
	row := dbManager.database.QueryRow("SELECT time FROM logs WHERE time = (SELECT MAX(time) FROM logs WHERE DATE(time) < DATE('now')) LIMIT 1")
	var t time.Time
	err := row.Scan(&t)
	return t, err
}

func (dbManager *DbManager) GetRecord(id int) (*LogRecord, error) {
	// subselect coz column affinity
	row := dbManager.database.QueryRow("SELECT id, message, time FROM logs WHERE id = ?", id)
	var r = LogRecord{}
	err := row.Scan(&r.Id, &r.Message, &r.Time)
	return &r, err
}

func (dbManager *DbManager) Set(id int, message string, messageTime time.Time) error {
	if !messageTime.Equal(time.Time{}) {
		stmt, err := dbManager.database.Prepare("UPDATE logs SET time = ? WHERE id = ?")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(messageTime, id)
		if err != nil {
			return err
		}
	}

	if message != "" {
		stmt, err := dbManager.database.Prepare("UPDATE logs SET message = ? WHERE id = ?")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(message, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dbManager *DbManager) Delete(id int) error {
	stmt, err := dbManager.database.Prepare("DELETE FROM logs WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
