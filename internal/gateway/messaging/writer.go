package messaging

import (
	"assessment-go-source-code-muhammad-aditya/internal/model"
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Writer[T model.Event] struct {
	Writer *kafka.Writer
	Topic  string
	Log    *logrus.Logger
}

func (w *Writer[T]) Write(ctx context.Context, event T) error {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		w.Log.WithError(err).Error("error marshaling event")
		return err
	}

	msg := kafka.Message{
		Topic: w.Topic,
		Key:   []byte(event.GetId()),
		Value: eventBytes,
	}

	err = w.Writer.WriteMessages(ctx, msg)
	if err != nil {
		w.Log.WithError(err).Error("error writing message")
		return err
	}

	return nil
}
