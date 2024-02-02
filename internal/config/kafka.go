package config

import (
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewKafkaReader(config *viper.Viper, log *logrus.Logger) *kafka.Reader {
	kafkaReaderConfig := kafka.ReaderConfig{
		Brokers: config.GetStringSlice("kafka.bootstrap.servers"),
		GroupID: config.GetString("kafka.group.id"),
	}

	reader := kafka.NewReader(kafkaReaderConfig)

	return reader
}

func NewKafkaWriter(config *viper.Viper, log *logrus.Logger) *kafka.Writer {
	kafkaWriterConfig := &kafka.WriterConfig{
		Brokers:  config.GetStringSlice("kafka.bootstrap.servers"),
		Balancer: &kafka.LeastBytes{},
	}

	writer := kafka.NewWriter(*kafkaWriterConfig)

	return writer
}
