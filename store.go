package main

type Store interface {
	GetAllSms() []Sms
	GetSms() Sms
}

type InMemoryStore struct {
}

func NewInMemoryStore() *InMemoryStore {
	return new(InMemoryStore)
}

func (store *InMemoryStore) GetAllSms() []Sms {
	return []Sms{}
}

func (store *InMemoryStore) GetSms() Sms {
	return Sms{}
}
