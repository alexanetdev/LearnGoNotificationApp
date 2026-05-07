package main

import (
	"LearnGoNotificationApp/enums"
	"LearnGoNotificationApp/models"
	"LearnGoNotificationApp/senders"
	"LearnGoNotificationApp/utils"
	"fmt"
)

func main() {
	fmt.Printf("TestApp is starting...\n")

	notifications := make(map[string]*models.Person)
	notifications["Alex"] = models.NewPerson("Alex", enums.Reminder, enums.NotSent, senders.SmsSender{})
	notifications["Chris"] = models.NewPerson("Chris", enums.Blast, enums.NotSent, senders.EmailSender{})
	notifications["Prince"] = models.NewPerson("Prince", enums.Blast, enums.NotSent, senders.SmsSender{})
	notifications["Albert"] = models.NewPerson("Albert", enums.Reminder, enums.NotSent, senders.EmailSender{})
	notifications["Kevin"] = models.NewPerson("Kevin", enums.Reminder, enums.NotSent, senders.SmsSender{})

	for key, value := range notifications {
		fmt.Printf("Try to send notification of type %s to %s\n", value.OutreachType, key)
		if value.OutreachType == enums.Reminder && value.Status != enums.Success {
			success, err := value.Sender.Send(key)

			if err != nil {
				fmt.Printf("\nCould not send reminder to %s: %s \n\n", key, err)
				notifications[key].Status = enums.Failed
			} else if success {
				fmt.Printf("Sent reminder to %s\n", key)
				notifications[key].Status = enums.Success
			}
		} else if value.OutreachType == enums.Blast {
			fmt.Printf("Outreach type of blast is not implemented for %s\n", key)
		} else {
			fmt.Printf("Not trying previously failed outreach")
		}
	}

	failed := utils.FilterMap(notifications, func(p *models.Person) bool {
		return p.Status == enums.Failed
	})

	successful := utils.FilterMap(notifications, func(p *models.Person) bool {
		return p.Status == enums.Success
	})

	remaining := utils.FilterMap(notifications, func(p *models.Person) bool {
		return p.Status == enums.NotSent
	})

	fmt.Printf("Failed Outreach:\n")
	fmt.Printf("Count: %d | Map: %v \n\n", len(failed), failed)

	fmt.Printf("Successful Outreach:\n")
	fmt.Printf("Count: %d | Map: %v \n\n", len(successful), successful)

	fmt.Printf("Remaining Outreach:\n")
	fmt.Printf("Count: %d | Map: %v \n\n", len(remaining), remaining)
}
