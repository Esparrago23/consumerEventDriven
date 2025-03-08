package notification

import (
	"bytes"
	"log"
	"net/http"
)


func SendNotification(notificationAPIURL string, message string) {
	jsonData := []byte(`{"message": "` + message + `"}`)

	resp, err := http.Post(notificationAPIURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error sending notification to API: %s", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Notification sent, received status: %s", resp.Status)
}
