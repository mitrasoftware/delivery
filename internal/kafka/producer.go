package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

var producer *kgo.Client

func InitProducer() {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(Brokers...),
		kgo.AllowAutoTopicCreation(),
	)

	if err != nil {
		log.Fatal("❌ Kafka producer init failed:", err)

	}

	producer = cl

	log.Println("✅ franz-go producer ready")

}

func Publish(topic string, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("marshal error:", err)
		return
	}

	record := &kgo.Record{
		Topic:     topic,
		Value:     data,
		Partition: 1,
	}

	producer.Produce(context.Background(), record, func(_ *kgo.Record, err error) {
		if err != nil {
			log.Println("❌ Kafka publish failed:", err)

		}
	})
}
