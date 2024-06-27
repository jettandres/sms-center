package main

import "strings"

type Store interface {
	GetAllSms() ([]Sms, error)
	GetAllSmsFromNumber(mobileNumber string) ([]Sms, error)
	GetSmsById(id string) (Sms, error)
	InsertSms(fromMobileNumber string, toMobileNumber string, body string) (Sms, error)
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

func (store *InMemoryStore) GetAllSmsFromNumber(mobileNumber string) ([]Sms, error) {
	allSmsFromNumber := make([]Sms, 0)

	for _, v := range store.SmsMessages {
		if strings.Compare(v.From, mobileNumber) == 0 {
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

func (store *InMemoryStore) InsertSms(fromMobileNumber string, toMobileNumber string, body string) (Sms, error) {
	sms := Sms{
		Id:          "some-new-id",
		Inserted_at: "06/21/24",
		From:        fromMobileNumber,
		To:          toMobileNumber,
		Body:        body,
	}
	store.SmsMessages = append(store.SmsMessages, sms)
	return sms, nil
}
