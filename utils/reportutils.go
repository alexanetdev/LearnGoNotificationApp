package utils

import (
	"LearnGoNotificationApp/models"
	"fmt"
)

func PrintStatusReport(label string, results map[string]*models.Person) {
	fmt.Printf("%s — Count: %d\n", label, len(results))
	for k, v := range results {
		fmt.Printf("  %s: %s\n", k, v)
	}
	fmt.Println()
}
