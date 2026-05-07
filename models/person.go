package models

import (
	"LearnGoNotificationApp/enums"
	"fmt"
)

type Person struct {
	Name         string
	OutreachType enums.OutreachType
	Status       enums.OutreachStatus
	Sender       NotificationSender
}

func NewPerson(name string, outreachType enums.OutreachType, status enums.OutreachStatus, sender NotificationSender) *Person {
	p := Person{Name: name, OutreachType: outreachType, Status: status, Sender: sender}

	return &p
}

func (p *Person) String() string {
	return fmt.Sprintf("%s | %s | %s", p.Name, p.OutreachType, p.Status)
}
