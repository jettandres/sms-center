package main

import (
	"fmt"
	"strings"
)

type Store interface {
	GetAllSms() ([]Sms, error)
	GetAllSmsFromSender(sender string) ([]Sms, error)
	GetSmsById(id string) (Sms, error)
	InsertSms(sender string, receiver string, body string) (Sms, error)
}

type InMemoryStore struct {
	SmsMessages []Sms
}

func NewInMemoryStore() *InMemoryStore {
	return new(InMemoryStore)
}

func (store *InMemoryStore) GetAllSms() ([]Sms, error) {
	return store.SmsMessages, nil
}

func (store *InMemoryStore) GetAllSmsFromSender(sender string) ([]Sms, error) {
	allSmsFromNumber := make([]Sms, 0)

	for _, v := range store.SmsMessages {
		if strings.Compare(v.Sender, sender) == 0 {
			allSmsFromNumber = append(allSmsFromNumber, v)
		}
	}
	return allSmsFromNumber, nil
}

func (store *InMemoryStore) GetSmsById(id string) (Sms, error) {
	for _, v := range store.SmsMessages {
		if v.Id == id {
			return v, nil
		}
	}
	return Sms{}, nil
}

func (store *InMemoryStore) InsertSms(sender string, receiver string, body string) (Sms, error) {
	id := fmt.Sprintf("some-new-id-%d", len(store.SmsMessages))
	sms := Sms{
		Id:          id,
		Inserted_at: "06/21/24",
		Sender:      sender,
		Receiver:    receiver,
		Body:        body,
	}
	store.SmsMessages = append(store.SmsMessages, sms)
	return sms, nil
}
