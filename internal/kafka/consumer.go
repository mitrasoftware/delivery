package kafka

import (
	"context"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

func StartConsumer() {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(Brokers...),
		kgo.ConsumerGroup("delivery-group"),
		kgo.ConsumeTopics("delivery.started"),
	)

	if err != nil {
		log.Fatal("❌ Kafka consumer init failed:", err)
	}

	go func() {

		for {
			fetches := cl.PollFetches(context.Background())
			if errs := fetches.Errors(); len(errs) > 0 {
				log.Println("Kafka fetch error:", errs)
				continue
			}

			fetches.EachRecord(func(record *kgo.Record) {

				log.Printf("📥 Kafka msg | topic=%s value=%s\n",
					record.Topic,
					string(record.Value),
				)
			})
		}
	}()
	log.Println("✅ franz-go consumer started")

}
