package main

import (
	"LearnGoNotificationApp/enums"
	"LearnGoNotificationApp/models"
	"LearnGoNotificationApp/senders"
	"LearnGoNotificationApp/utils"
	"LearnGoNotificationApp/workers"
	"fmt"
	"sync"
)

func main() {
	fmt.Printf("TestApp is starting...\n")
	fmt.Println()

	notifications := make(map[string]*models.Person)
	notifications["Alex"] = models.NewPerson("Alex", enums.Reminder, enums.NotSent, senders.SmsSender{})
	notifications["Chris"] = models.NewPerson("Chris", enums.Blast, enums.NotSent, senders.EmailSender{})
	notifications["Prince"] = models.NewPerson("Prince", enums.Blast, enums.NotSent, senders.SmsSender{})
	notifications["Albert"] = models.NewPerson("Albert", enums.Reminder, enums.NotSent, senders.EmailSender{})
	notifications["Kevin"] = models.NewPerson("Kevin", enums.Reminder, enums.NotSent, senders.SmsSender{})

	notificationQueue := make(chan *models.Person)

	numWorkers := 3

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workers.NotificationWorker(notificationQueue)
		}()
	}

	// fill the channel
	for key, value := range notifications {
		fmt.Printf("Queueing notification for %s\n", key)
		notificationQueue <- value
	}

	close(notificationQueue)
	wg.Wait()

	failed := utils.FilterMap(notifications, func(p *models.Person) bool {
		return p.Status == enums.Failed
	})

	successful := utils.FilterMap(notifications, func(p *models.Person) bool {
		return p.Status == enums.Success
	})

	remaining := utils.FilterMap(notifications, func(p *models.Person) bool {
		return p.Status == enums.NotSent
	})

	fmt.Println()

	utils.PrintStatusReport("Failed Outreach", failed)
	utils.PrintStatusReport("Successful Outreach", successful)
	utils.PrintStatusReport("Remaining Outreach", remaining)
}
