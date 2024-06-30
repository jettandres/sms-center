package main

type Store interface {
	GetAllSms() ([]Sms, error)
	GetAllSmsFromSender(sender string) ([]Sms, error)
	GetSmsById(id string) (Sms, error)
	InsertSms(sender string, receiver string, body string) (Sms, error)
}
