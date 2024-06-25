package main

type Store interface {
	GetAllSms() []Sms
}

type InMemoryStore struct {
}

func (store *InMemoryStore) GetAllSms() []Sms {
	return []Sms{}
}

func NewInMemoryStore() *InMemoryStore {
	return new(InMemoryStore)
}
