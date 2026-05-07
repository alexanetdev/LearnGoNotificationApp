package models

type NotificationSender interface {
	Send(name string) (bool, error)
}
