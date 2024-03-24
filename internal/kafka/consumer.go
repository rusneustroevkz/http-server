package kafka

import (
	"github.com/rusneustroevkz/http-server/pkg/logger"
	"log"

	"github.com/IBM/sarama"
)

type Consumer struct {
	Handle func(data []byte) error
	log    logger.Logger
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	defer func() {
		if recoverMsg := recover(); recoverMsg != nil {
			c.log.Error("has panic", logger.Any("message", recoverMsg))
		}
	}()
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}

			if err := c.Handle(message.Value); err != nil {
				c.log.Error("cannot handle data", logger.String("topic", message.Topic), logger.Error(err))
			}

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
