package messaging

import (
	"assessment-go-source-code-muhammad-aditya/internal/model"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type CustomerWriter struct {
	Writer[*model.CustomerEvent]
}

func NewCustomerWriter(writer *kafka.Writer, log *logrus.Logger) *CustomerWriter {
	return &CustomerWriter{
		Writer: Writer[*model.CustomerEvent]{
			Writer: writer,
			Topic:  "customers",
			Log:    log,
		},
	}
}
