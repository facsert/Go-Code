package database

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"strings"

	_ "modernc.org/sqlite"

	"learn/comm"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(dir string) *sql.DB {
	db, err := sql.Open("sqlite", dir)
	if err != nil {
		slog.Error(fmt.Sprintf("Open DB file error: %v", err))
		log.Fatal(err)
	}

	return db
}

func (d *Database) Query(sql string, args ...any) (*sql.Rows, error) {
	return d.db.Query(sql, args...)
}

func (d *Database) Exec(sql string, args ...any) (sql.Result, error) {
	return d.db.Exec(sql, args...)
}

type Record struct {
	File    string
	Modify  string
	Created string
}

func CreateTable(db *sql.DB, table string) error {
	sql := `CREATE TABLE IF NOT EXISTS %s (
        file TEXT NOT NULL,
        modify DATETIME NOT NULL,
        created DATE NOT NULL,
        CONSTRAINT unique_data UNIQUE(file, modify)
	);`

	slog.Info(fmt.Sprintf(sql, table))
	_, err := db.Exec(fmt.Sprintf(sql, table))

	if err != nil {
		return err
	}
	return nil
}

func InsertData(db *sql.DB, table string, records []Record) error {
	sql := "INSERT INTO %s (file, modify, created) VALUES \n%s \nON CONFLICT(file, modify)  DO NOTHING;"

	lines := []string{}
	for _, r := range records {
		lines = append(lines, fmt.Sprintf("('%s', '%s', '%s')", r.File, r.Modify, r.Created))
	}
	slog.Info(fmt.Sprintf(sql, table, strings.Join(lines, ",\n")))

	_, err := db.Exec(fmt.Sprintf(sql, table, strings.Join(lines, ",")))
	if err != nil {
		return err
	}
	return nil
}

func QueryData(db *sql.DB, table string, created string) ([]Record, error) {
	sql := `SELECT file, modify, created FROM %s WHERE created = '%s';`
	slog.Info(fmt.Sprintf(sql, table, created))

	rows, err := db.Query(fmt.Sprintf(sql, table, created))
	if err != nil {
		return []Record{}, err
	}
	var records []Record
	for rows.Next() {
		var r Record
		err = rows.Scan(&r.File, &r.Modify, &r.Created)
		if err != nil {
		}
		records = append(records, r)
	}
	return records, nil
}

// go run main.go
func SqliteTest() {
	var DATABASE_DIR = comm.AbsPath("file.db")
	var DATABASE_TABLE = "record"

	cur := NewDatabase(DATABASE_DIR)
	err := CreateTable(cur, DATABASE_TABLE)
	if err != nil {
		slog.Error(fmt.Sprintf("Create Database error: %v\n", err))
		return
	}

	rs := []Record{
		{File: "file0.txt", Modify: "2023-01-01 08:08:08", Created: "2023-01-01"},
		{File: "file1.txt", Modify: "2023-01-02 08:08:09", Created: "2023-01-01"},
		{File: "file2.txt", Modify: "2023-01-03 08:08:10", Created: "2023-01-01"},
	}
	err = InsertData(cur, DATABASE_TABLE, rs)
	if err != nil {
		slog.Error(fmt.Sprintf("Insert Table error: %v\n", err))
		return
	}

	list, err := QueryData(cur, DATABASE_TABLE, "2023-01-01")
	if err != nil {
		slog.Error(fmt.Sprintf("Query Table error: %v\n", err))
		return
	}

	for _, line := range list {
		slog.Info(fmt.Sprintf("%#v\n", line))
	}
}
