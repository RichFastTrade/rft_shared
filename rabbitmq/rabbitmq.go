package rabbitmq

import (
	"fmt"
	"github.com/RichFastTrade/rft_shared/conf"
	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
)

var conn *amqp.Connection
var onceConn sync.Once
var channel *amqp.Channel
var onceChannel sync.Once

func getConn() (*amqp.Connection, error) {
	var connErr error
	onceConn.Do(func() {
		connStr := fmt.Sprintf(
			"amqp://%s:%s@%s:%d/",
			conf.Get().Rabbitmq.User,
			conf.Get().Rabbitmq.Password,
			conf.Get().Rabbitmq.Host,
			conf.Get().Rabbitmq.Port,
		)
		log.Infof("connecting to %s", connStr)
		c, err := amqp.Dial(connStr)
		if err != nil {
			connErr = err
			return
		}
		conn = c
	})
	return conn, connErr
}

func getChannel() (*amqp.Channel, error) {
	var channelErr error
	onceChannel.Do(func() {
		conn, err := getConn()
		if err != nil {
			channelErr = err
			return
		}
		c, err := conn.Channel()
		if err != nil {
			channelErr = err
			return
		}
		channel = c
	})
	return channel, channelErr
}

func RegisterQueue(name string) error {
	c, err := getChannel()
	if err != nil {
		return err
	}
	_, err = c.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	return err
}

func RegisterQueues(names []string) error {
	for _, name := range names {
		if err := RegisterQueue(name); err != nil {
			return err
		}
	}
	return nil
}

func Publish(queue string, body []byte) error {
	ch, err := getChannel()
	if err != nil {
		return err
	}
	if err := ch.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	); err != nil {
		return err
	}
	return nil
}

func consume(queue string) (<-chan amqp.Delivery, error) {
	ch, err := getChannel()
	if err != nil {
		return nil, err
	}
	return ch.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}

func RegisterConsume(queue string, operation func([]byte) error) error {
	msgs, err := consume(queue)
	if err != nil {
		return err
	}
	go func() {
		for msg := range msgs {
			if err := operation(msg.Body); err != nil {
				log.Errorf("Failed to process message: %s", err)
				if err := msg.Reject(true); err != nil {
					log.Errorf("Failed to reject message: %s", err)
				}
				continue
			}
			if err := msg.Ack(false); err != nil {
				log.Errorf("Failed to ack message: %s", err)
			}
		}
	}()
	return nil
}
