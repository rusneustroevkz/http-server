package config

import (
	"github.com/rusneustroevkz/http-server/pkg/logger"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	GRPCServer    Server `yaml:"grpc-server"`
	PublicServer  Server `yaml:"public-server"`
	PrivateServer Server `yaml:"private-server"`
	Kafka         Kafka  `yaml:"kafka"`
	App           App    `yaml:"app"`
}

type App struct {
	Production        bool `yaml:"production"`
	RequestLogEnabled bool `yaml:"request-log-enabled"`
}

type Server struct {
	Port    int64 `yaml:"port"`
	Timeout int64 `yaml:"timeout"`
}

type Kafka struct {
	Verbose    bool       `yaml:"verbose"`
	ClientName string     `yaml:"client-name"`
	Brokers    []string   `yaml:"brokers"`
	Consumers  []Consumer `yaml:"consumers"`
	Producers  []Producer `yaml:"producers"`
}

type Consumer struct {
	Name          string         `yaml:"name"`
	ConsumerGroup string         `yaml:"consumer-group"`
	Topics        []string       `yaml:"topics"`
	Config        ConsumerConfig `yaml:"config"`
}

type Producer struct {
	Name   string         `yaml:"name"`
	Topic  string         `yaml:"topics"`
	Config ProducerConfig `yaml:"config"`
}

type ConsumerConfig struct {
}

type ProducerConfig struct {
}

func NewConfig(log logger.Logger) *Config {
	var cfg Config
	fp, err := filepath.Abs("config/config.yaml")
	if err != nil {
		log.Fatal("cant get filepath", logger.Error(err))
	}
	configFile, err := os.ReadFile(fp)
	if err != nil {
		log.Fatal("cant read config file", logger.Error(err))
	}
	if err := yaml.Unmarshal(configFile, &cfg); err != nil {
		log.Fatal("cant unmarshal config", logger.Error(err))
	}
	return &cfg
}
