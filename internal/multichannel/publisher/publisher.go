package multichannel

import (
	"fmt"
	"log"
	"sync"

	rabbitmq "github.com/hadihammurabi/go-rabbitmq"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	def, err := rabbitmq.New("ampq://guest:guest@localhost:5672/")
	failOnError(err, fmt.Sprintf("%v", err))
	defer def.Close()

	del, err := rabbitmq.NewFromConnection(def.Connection())
	failOnError(err, fmt.Sprintf("%v", err))
	defer del.Close()

	_, err = del.Queue().
		WithName("hello").
		WithChannel(del.Channel()).
		Declare()

	failOnError(err, fmt.Sprintf("%v", err))

	var wg sync.WaitGroup
	max := 10
	wg.Add(max)
	for i := 0; i < max; i++ {
		go func(a int) {
			defer wg.Done()
			body := fmt.Sprintf("Hello World %d !", a)
			err = del.Publish(&rabbitmq.MQConfigPublish{
				RoutingKey: del.Queue().Name,
				Message: amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				},
			})
		}(i)
	}
	wg.Wait()
}
