package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteStore struct {
	driverName string
	dsn        string
	DB         *sqlx.DB
}

func NewSqliteStore() *SqliteStore {
	store := &SqliteStore{
		driverName: "sqlite3",
		dsn:        os.Getenv("SQL_SQLITE_DSN"),
	}

	db, err := sqlx.Open(store.driverName, store.dsn)
	if err != nil {
		panic(err.Error())
	}
	store.DB = db

	return store
}

func (store *SqliteStore) GetAllSms() ([]Sms, error) {
	allSms := make([]Sms, 0)

	stmt, err := store.DB.Preparex("SELECT * FROM sms_messages")
	defer stmt.Close()

	if err != nil {
		fmt.Printf("[SQLITE_STORE][STATEMENT] Error: %s", err.Error())
		panic(err.Error())
	}

	err = stmt.Select(&allSms)

	if err != nil {
		fmt.Printf("[SQLITE_STORE][QUERY] Error: %s", err.Error())
		panic(err.Error())
	}

	return allSms, nil
}

func (store *SqliteStore) GetAllSmsFromSender(sender string) ([]Sms, error) {
	allSmsFromSender := make([]Sms, 0)

	stmt, err := store.DB.Preparex("SELECT * FROM sms_messages WHERE sender = ?")
	defer stmt.Close()

	if err != nil {
		fmt.Printf("[SQLITE_STORE][STATEMENT] Error: %s", err.Error())
		panic(err.Error())
	}

	err = stmt.Select(&allSmsFromSender, sender)

	if err != nil {
		fmt.Printf("[SQLITE_STORE][QUERY] Error: %s", err.Error())
		panic(err.Error())
	}

	return allSmsFromSender, nil
}

func (store *SqliteStore) GetSmsById(id string) (Sms, error) {
	stmt, err := store.DB.Preparex("SELECT * FROM sms_messages WHERE id = ?")
	defer stmt.Close()

	if err != nil {
		fmt.Printf("[SQLITE_STORE][STATEMENT] Error: %s", err.Error())
		panic(err.Error())
	}

	var sms Sms
	err = stmt.Get(&sms, id)

	if err != nil {
		fmt.Printf("[GET_SMS_BY_ID][QUERY] Error: %s", err.Error())
		panic(err.Error())
	}

	return sms, nil
}

func (store *SqliteStore) InsertSms(sender string, receiver string, body string) (Sms, error) {
	stmt, err := store.DB.Preparex("INSERT INTO sms_messages (id, inserted_at, body, sender, receiver) VALUES (?,datetime(),?,?,?)")
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
