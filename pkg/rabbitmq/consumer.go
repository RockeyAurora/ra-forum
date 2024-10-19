package rabbitmq

import (
	"bluebell/models"
	"bluebell/pkg/email"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

func Consumer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"email_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var Ed models.ParamEmailData
			err := json.Unmarshal(d.Body, &Ed)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}
			log.Printf("Sending email for user: %s", Ed.Username)
			// 处理邮件发送逻辑
			err = email.SendEmail(&Ed)
			if err != nil {
				log.Printf("Error sending email: %s", err)
			}
		}
	}()

	<-forever
}
