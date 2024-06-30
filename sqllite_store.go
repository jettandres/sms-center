package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteStore struct {
	driverName string
	dsn        string
	DB         *sql.DB
}

func NewSqliteStore() *SqliteStore {
	store := &SqliteStore{
		driverName: "sqlite3",
		dsn:        os.Getenv("SQL_SQLITE_DSN"),
	}

	db, err := sql.Open(store.driverName, store.dsn)
	if err != nil {
		panic(err.Error())
	}
	store.DB = db

	return store
}

func (store *SqliteStore) GetAllSms() ([]Sms, error) {
	allSms := make([]Sms, 0)

	stmt, err := store.DB.Prepare("SELECT id, body, sender, receiver, inserted_at FROM sms_messages")
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
	allSmsFromSender := make([]Sms, 0)

	stmt, err := store.DB.Prepare("SELECT id, body, sender, receiver, inserted_at FROM sms_messages WHERE sender = ?")
	defer stmt.Close()

	if err != nil {
		fmt.Printf("[SQLITE_STORE][STATEMENT] Error: %s", err.Error())
		panic(err.Error())
	}

	rows, err := stmt.Query(sender)
	defer rows.Close()

	if err != nil {
		fmt.Printf("[SQLITE_STORE][QUERY] Error: %s", err.Error())
		panic(err.Error())
	}

	for rows.Next() {
		var sms Sms
		if err := rows.Scan(&sms.Id, &sms.Body, &sms.Sender, &sms.Receiver, &sms.Inserted_at); err != nil {
			return nil, fmt.Errorf("[GET_ALL_SMS_FROM_SENDER][ERROR] %s", err.Error())
		}
		allSmsFromSender = append(allSmsFromSender, sms)
	}

	return allSmsFromSender, nil
}

func (store *SqliteStore) GetSmsById(id string) (Sms, error) {
	stmt, err := store.DB.Prepare("SELECT id, body, sender, receiver, inserted_at FROM sms_messages WHERE id = ?")
	defer stmt.Close()

	if err != nil {
		fmt.Printf("[SQLITE_STORE][STATEMENT] Error: %s", err.Error())
		panic(err.Error())
	}

	var sms Sms
	err = stmt.QueryRow(id).Scan(&sms.Id, &sms.Body, &sms.Sender, &sms.Receiver, &sms.Inserted_at)

	if err != nil {
		fmt.Printf("[GET_SMS_BY_ID][QUERY] Error: %s", err.Error())
		panic(err.Error())
	}

	return sms, nil
}

func (store *SqliteStore) InsertSms(sender string, receiver string, body string) (Sms, error) {
	stmt, err := store.DB.Prepare("INSERT INTO sms_messages (id, inserted_at, body, sender, receiver) VALUES (?,datetime(),?,?,?)")
	defer stmt.Close()

	if err != nil {
		fmt.Printf("[SQLITE_STORE][STATEMENT] Error: %s", err.Error())
		panic(err.Error())
	}

	sms := Sms{
		Id:       uuid.NewString(),
		Body:     body,
		Sender:   sender,
		Receiver: receiver,
	}

	res, err := stmt.Exec(sms.Id, sms.Body, sms.Sender, sms.Receiver)
	affected, err := res.RowsAffected()

	if affected == 0 {
		fmt.Printf("[SQLITE_STORE][INSERT] Error: %s", err.Error())
		return Sms{}, err
	}

	return sms, nil
}
