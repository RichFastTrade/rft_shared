package initial

import "github.com/RichFastTrade/rft_shared/rabbitmq"

func initRabbitMQ() error {
	if err := rabbitmq.RegisterQueues([]string{
		"update",
		"message",
		"analyze",
	}); err != nil {
		return err
	}
	return nil
}
