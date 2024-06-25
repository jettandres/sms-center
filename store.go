package main

type Store interface {
	GetAllSms() []Sms
	GetSms(mobilenNumber string) Sms
}

type InMemoryStore struct {
}

func NewInMemoryStore() *InMemoryStore {
	return new(InMemoryStore)
}

func (store *InMemoryStore) GetAllSms() []Sms {
	return []Sms{}
}

func (store *InMemoryStore) GetSms(mobileNumber string) Sms {
	return Sms{}
}
