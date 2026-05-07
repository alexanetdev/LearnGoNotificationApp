package senders

import (
	"fmt"
)

type SmsSender struct{}

func (s SmsSender) Send(name string) (bool, error) {
	if name == "Alex" {
		return false, fmt.Errorf("Name cannot be Alex")
	}

	return true, nil
}
