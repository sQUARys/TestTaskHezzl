package kafka

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
	"sync"
)

const (
	topic         = "new_topic"
	brokerAddress = "192.168.1.103:9092"
)

type Kafka struct {
	writer *kafka.Writer
	reader *kafka.Reader
}

func New() *Kafka {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:     []string{brokerAddress},
		Topic:       topic,
		Logger:      kafka.LoggerFunc(logf),
		ErrorLogger: kafka.LoggerFunc(logf),
	})
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{brokerAddress},
		Topic:       topic,
		GroupID:     "test-consumer-group",
		Partition:   0,
		Logger:      kafka.LoggerFunc(logf),
		ErrorLogger: kafka.LoggerFunc(logf),
	})

	return &Kafka{
		writer: w,
		reader: r,
	}
}

func (k *Kafka) Start(wg sync.WaitGroup) {
	defer wg.Done()

	ctx := context.Background()

	k.WriteText("Hello from kafka", ctx)
	k.ReadText(ctx)
}

func (k *Kafka) ReadText(ctx context.Context) error {

	msg, err := k.reader.ReadMessage(ctx)
	if err != nil {
		return err
	}

	fmt.Println("received: ", string(msg.Value))
	return nil
}

func (k *Kafka) WriteText(log string, ctx context.Context) error {

	err := k.writer.WriteMessages(ctx, kafka.Message{
		Value: []byte(log),
	})
	if err != nil {
		return err
	}
	return nil
}

func logf(msg string, a ...interface{}) {
	fmt.Printf("Log Kafka :"+msg, a...)
	fmt.Println()
}
