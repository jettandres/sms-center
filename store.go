package main

type Store interface {
	GetAllSms() []Sms
	GetAllSmsFromNumber(mobileNumber string) []Sms
	GetSmsFromNumber(mobileNumber string) Sms
}

type InMemoryStore struct {
}

func NewInMemoryStore() *InMemoryStore {
	return new(InMemoryStore)
}

func (store *InMemoryStore) GetAllSms() []Sms {
	return []Sms{}
}

func (store *InMemoryStore) GetAllSmsFromNumber(mobileNumber string) []Sms {
	return []Sms{}
}

func (store *InMemoryStore) GetSmsFromNumber(mobileNumber string) Sms {
	return Sms{}
}
