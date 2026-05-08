package workers

import (
	"LearnGoNotificationApp/enums"
	"LearnGoNotificationApp/models"
	"fmt"
)

func NotificationWorker(queue <-chan *models.Person) {
	for person := range queue {
		processNotification(person)
	}
}

func processNotification(person *models.Person) {
	fmt.Printf("Try to send notification of type %s to %s\n", person.OutreachType, person.Name)
	if person.OutreachType == enums.Reminder && person.Status != enums.Success {
		success, err := person.Sender.Send(person.Name)
		if err != nil {
			fmt.Printf("Could not send reminder to %s: %s \n", person.Name, err)
			person.Status = enums.Failed
		} else if success {
			fmt.Printf("Sent reminder to %s\n", person.Name)
			person.Status = enums.Success
		}
	} else if person.OutreachType == enums.Blast {
		fmt.Printf("Outreach type of blast is not implemented for %s\n", person.Name)
	} else {
		fmt.Printf("Not trying previously failed outreach")
	}
}
