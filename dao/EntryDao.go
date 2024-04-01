package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"omega-audit-service/config"
	"omega-audit-service/models"
	"omega-audit-service/util"
	"time"
)

// AddEntry /*
func AddEntry(entry models.Entry) bool {
	fmt.Println("Adding entry : " + entry.SourceUrl)
	var success = true
	db := connect()
	insert, err := db.Query(util.GetInsertQuery(),
		entry.Source, entry.SourceUrl, entry.ResourceUrl,
		entry.SystemFilePath,
		entry.Strategy,
		entry.Invoker,
		entry.Channel,
		entry.Title, time.Now())

	if err != nil {
		_ = fmt.Errorf(err.Error())
		success = false
	}
	defer db.Close()
	defer insert.Close()

	return success
}

// GetEntriesBySourceUrl /*
func GetEntriesBySourceUrl(url string) ([]models.Entry, error) {
	fmt.Println("Requesting search for " + url)
	db := connect()
	rows, err := db.Query(util.GetEntriesQuery(url))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	defer db.Close()
	var entries = make([]models.Entry, 0)
	for rows.Next() {
		var entry models.Entry
		if err := rows.Scan(&entry.Source, &entry.SourceUrl,
			&entry.ResourceUrl,
			&entry.SystemFilePath,
			&entry.Strategy,
			&entry.Invoker,
			&entry.Channel,
			&entry.Title,
			&entry.Timestamp); err != nil {
			return entries, err
		}
		entries = append(entries, entry)
	}
	if err = rows.Err(); err != nil {
		return entries, err
	}
	return entries, nil
}

func connect() *sql.DB {
	db, err := sql.Open(config.GetDriverName(), config.GetDataSourceConfig())
	if err != nil {
		panic(err.Error())
	}
	return db
}
