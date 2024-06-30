package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type SqliteStore struct {
	driverName string
	dsn        string
}

func NewSqliteStore() *SqliteStore {
	store := &SqliteStore{
		driverName: "sqlite3",
		dsn:        os.Getenv("SQL_SQLITE_DSN"),
	}

	return store
}

func (store *SqliteStore) Database() *sql.DB {
	db, err := sql.Open(store.driverName, store.dsn)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func (store *SqliteStore) GetAllSms() ([]Sms, error) {
	allSms := make([]Sms, 0)
	defer store.Database().Close()

	stmt, err := store.Database().Prepare("SELECT * FROM sms_messages")
	defer stmt.Close()

	if err != nil {
		fmt.Printf("[SQLITE_STORE][STATEMENT] Error: %s", err.Error())
		panic(err.Error())
	}

	rows, err := stmt.Query()
	defer rows.Close()

	if err != nil {
		fmt.Printf("[SQLITE_STORE][QUERY] Error: %s", err.Error())
		panic(err.Error())
	}

	for rows.Next() {
		var sms Sms
		if err := rows.Scan(&sms.Id, &sms.Body, &sms.Sender, &sms.Receiver, &sms.Inserted_at); err != nil {
			return nil, fmt.Errorf("[GET_ALL_SMS][ERROR] %s", err.Error())
		}
		allSms = append(allSms, sms)
	}

	return allSms, nil
}

func (store *SqliteStore) GetAllSmsFromSender(sender string) ([]Sms, error) {
	return []Sms{}, nil
}

func (store *SqliteStore) GetSmsById(id string) (Sms, error) {
	return Sms{}, nil
}

func (store *SqliteStore) InsertSms(sender string, receiver string, body string) (Sms, error) {
	return Sms{}, nil
}
