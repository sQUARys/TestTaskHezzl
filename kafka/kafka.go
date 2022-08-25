package kafka

import (
	"github.com/segmentio/kafka-go"
)

const (
	topic         = "message_log"
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

//CREATE TABLE logs_try (
//logs_id Int32 Codec(DoubleDelta, LZ4),
//message String
//) Engine = MergeTree
//PRIMARY KEY (logs_id);

//CREATE TABLE log_try_queue (
//readings_id Int32,
//message String
//)
//ENGINE = Kafka
//SETTINGS kafka_broker_list = '192.168.1.103:9092',
//kafka_topic_list = 'message_log',
//kafka_group_name = 'my-group',
//kafka_format = 'SQLInsert';

type Kafka struct {
	Writer *kafka.Writer
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
		Writer: w,
		reader: r,
	}

}

//
//func (k *Kafka) ReadLog(ctx context.Context) {
//
//	msg, err := k.reader.ReadMessage(ctx)
//	if err != nil {
//		panic("could not read message " + err.Error())
//	}
//
//	fmt.Println("received: ", string(msg.Value))
//}
//
//func (k *Kafka) WriteLog(key string, log string, ctx context.Context) {
//
//	err := k.writer.WriteMessages(ctx, kafka.Message{
//		Key: []byte(key),
//		// create an arbitrary message payload for the value
//		Value: []byte(log),
//	})
//	if err != nil {
//		panic("could not write message " + err.Error())
//	}
//}
