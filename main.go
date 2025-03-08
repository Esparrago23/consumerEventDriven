package main

import (
	"log"
	"os"
	"consumer_Event_Driven/src/consumer"
	"consumer_Event_Driven/src/notification"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	rabbitmqQueue := os.Getenv("RABBITMQ_QUEUE")
	notificationAPIURL := os.Getenv("NOTIFICATION_API_URL")


	conn, ch := consumer.ConnectToRabbitMQ(rabbitmqURL)
	defer conn.Close()
	defer ch.Close()

	q := consumer.DeclareQueue(ch, rabbitmqQueue)


	msgs := consumer.StartMessageConsumer(ch, q)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			notification.SendNotification(notificationAPIURL, string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
