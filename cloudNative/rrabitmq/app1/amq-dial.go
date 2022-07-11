import (
	"os"
	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic("could not establish AMQP connection: " + err.Error())
	}
	channel, err := connection.Channel()
	if err != nil {
		panic("could not open channel: " + err.Error())
	}
	message := amqp.Publishing{
		Body: []byte("Hello World"),
	}
	err = channel.Publish("events", "some-routing-key", false, false, message)
	if err != nil {
		panic("error while publishing message: " + err.Error())
	}
	_, err = channel.QueueDeclare("my_queue", true, false, false, false, nil)
	if err != nil {
		panic("error while declaring the queue: " + err.Error())
	}

	err = channel.QueueBind("my_queue", "#", "events", false, nil)
	if err != nil {
		panic("error while binding the queue: " + err.Error())
	}
	msgs, err := channel.Consume("my_queue", "", false, false, false, false, nil)
	if err != nil {
		panic("error while consuming the queue: " + err.Error())
	}
	for msg := range msgs {
		fmt.Println("message received: " + string(msg.Body))
		msg.Ack(false)
	}
	defer connection.Close()
} 