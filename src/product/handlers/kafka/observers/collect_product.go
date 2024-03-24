package observers

import (
	"github.com/rusneustroevkz/http-server/internal/config"
	"github.com/rusneustroevkz/http-server/pkg/logger"
)

const consumerCollectProduct = "test-consumer-name"

type CollectProduct struct {
	cfg *config.Config
	log logger.Logger
}

func NewCollectProduct(cfg *config.Config, log logger.Logger) *CollectProduct {
	return &CollectProduct{
		cfg: cfg,
		log: log,
	}
}

func (c *CollectProduct) Handle(data []byte) error {
	c.log.Info("handle topic", logger.String("data", string(data)))
	return nil
}

func (c *CollectProduct) Consumer() *config.Consumer {
	for _, consumer := range c.cfg.Kafka.Consumers {
		if consumer.Name == consumerCollectProduct {
			return &consumer
		}
	}

	return nil
}
