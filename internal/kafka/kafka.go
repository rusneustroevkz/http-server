package kafka

import (
	"context"
	"errors"

	"github.com/rusneustroevkz/http-server/internal/config"
	"github.com/rusneustroevkz/http-server/pkg/logger"

	"github.com/IBM/sarama"
)

type Client struct {
	saramaConfig   *sarama.Config
	cfg            *config.Config
	log            logger.Logger
	consumerGroups []sarama.ConsumerGroup
	keepRunning    bool
}

func NewClient(cfg *config.Config, log logger.Logger) *Client {
	saramaConfig := newSaramaConfig(cfg, log)
	if cfg.Kafka.Verbose {
		sarama.Logger = log.Std()
	}

	return &Client{
		saramaConfig: saramaConfig,
		cfg:          cfg,
		log:          log,
		keepRunning:  true,
	}
}

func (k *Client) Run(_ context.Context) error {
	for _, c := range k.cfg.Kafka.Consumers {
		c := c
		consumer := Consumer{}

		go func() {
			ctx := context.Background()
			consumerGroup, err := sarama.NewConsumerGroup(k.cfg.Kafka.Brokers, c.ConsumerGroup, k.saramaConfig)
			if err != nil {
				k.log.Fatal("Error creating consumer group consumerGroup", logger.Error(err))
			}
			k.consumerGroups = append(k.consumerGroups, consumerGroup)
			k.log.Info("topic started", logger.String("name", c.Name))

			for k.keepRunning {
				if err := consumerGroup.Consume(ctx, c.Topics, &consumer); err != nil {
					if errors.Is(err, sarama.ErrClosedConsumerGroup) {
						return
					}
					k.log.Error("Error from consumer", logger.Error(err))
				}

				if ctx.Err() != nil {
					return
				}
			}
		}()
	}

	return nil
}

func (k *Client) Stop(_ context.Context) error {
	k.keepRunning = false
	for _, c := range k.consumerGroups {
		if err := c.Close(); err != nil {
			k.log.Fatal("Error closing client", logger.Error(err))
		}
	}
	return nil
}

func newSaramaConfig(cfg *config.Config, log logger.Logger) *sarama.Config {
	saramaConfig := sarama.NewConfig()
	version, err := sarama.ParseKafkaVersion(sarama.DefaultVersion.String())
	if err != nil {
		log.Fatal("cant parse kafka version", logger.Error(err))
	}

	saramaConfig.ClientID = cfg.Kafka.ClientName
	saramaConfig.Version = version

	if err := saramaConfig.Validate(); err != nil {
		log.Fatal("sarama config validate failed", logger.Error(err))
	}

	return saramaConfig
}
