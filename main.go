package main

import (
	"LearnGoNotificationApp/enums"
	"LearnGoNotificationApp/models"
	"LearnGoNotificationApp/senders"
	"fmt"
)

func main() {
	fmt.Printf("TestApp is starting...\n")

	notifications := make(map[string]*models.Person)
	notifications["Alex"] = models.NewPerson("Alex", enums.Reminder, "NotSent", senders.SmsSender{})
	notifications["Chris"] = models.NewPerson("Chris", enums.Blast, "NotSent", senders.EmailSender{})
	notifications["Prince"] = models.NewPerson("Prince", enums.Blast, "NotSent", senders.SmsSender{})
	notifications["Albert"] = models.NewPerson("Albert", enums.Reminder, "NotSent", senders.EmailSender{})
	notifications["Kevin"] = models.NewPerson("Kevin", enums.Reminder, "NotSent", senders.SmsSender{})

	for key, value := range notifications {
		fmt.Printf("Try to send notification of type %s to %s\n", value.OutreachType, key)
		if value.OutreachType == enums.Reminder && value.Status != "Success" {
			success, err := value.Sender.Send(key)

			if err != nil {
				fmt.Printf("\nCould not send reminder to %s: %s \n\n", key, err)
				notifications[key].Status = "Failed"
			} else if success {
				fmt.Printf("Sent reminder to %s\n", key)
				notifications[key].Status = "Success"
			}
		} else if value.OutreachType == enums.Blast {
			fmt.Printf("Outreach type of blast is not implemented for %s\n", key)
		} else {
			fmt.Printf("Not trying previously failed outreach")
		}
	}

	fmt.Printf("Notification Statuses: \n")
	for k, v := range notifications {
		fmt.Printf("%s : %v\n", k, v)
	}
}
