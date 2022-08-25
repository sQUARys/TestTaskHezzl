package kafka

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
)

const (
	topic         = "new_topic"
	brokerAddress = "192.168.1.103:9092"
)

//CREATE TABLE log_queue (
//readings_id Int32,
//time DateTime,
//temperature Decimal(5,2)
//)
//ENGINE = Kafka
//SETTINGS kafka_broker_list = '192.168.1.103:9092',
//kafka_topic_list = 'message_log',
//kafka_group_name = 'my-group',
//kafka_format = 'SQLInsert';

//CREATE TABLE logs (
//readings_id Int32 Codec(DoubleDelta, LZ4),
//time DateTime Codec(DoubleDelta, LZ4),
//date ALIAS toDate(time),
//temperature Decimal(5,2) Codec(T64, LZ4)
//) Engine = MergeTree
//PARTITION BY toYYYYMM(time)
//ORDER BY (readings_id, time);
//
//CREATE TABLE log2 (
//value String,
//) Engine = MergeTree
//PRIMARY KEY (value);

//CREATE TABLE logsSec_queue (
//value String
//)
//ENGINE = Kafka
//SETTINGS kafka_broker_list = '192.168.1.103:9092',
//kafka_topic_list = 'new_topic',
//kafka_group_name = 'my-group',
//kafka_format = 'Vertical';

type Kafka struct {
	writer *kafka.Writer
	reader *kafka.Reader
}

func New() *Kafka {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
	})

	return &Kafka{
		writer: w,
		reader: r,
	}
}

func (k *Kafka) ReadLog(ctx context.Context) error {

	msg, err := k.reader.ReadMessage(ctx)
	if err != nil {
		return err
	}

	fmt.Println("received: ", string(msg.Value))
	return nil
}

func (k *Kafka) WriteLog(log string, ctx context.Context) error {

	err := k.writer.WriteMessages(ctx, kafka.Message{
		Value: []byte(log),
	})
	if err != nil {
		return err
	}
	return nil
}
