package franz

import (
	"context"
	"time"

	cfg "github.com/spf13/viper"
	"github.com/twmb/franz-go/pkg/kgo"
)

// NewKafkaClient Return new Kafka client instance
func NewKafkaClient(ctx context.Context) (client *kgo.Client, err error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(cfg.GetStringSlice("Kafka.BROKERS")...),
		kgo.ClientID(cfg.GetString("Kafka.CLIENT_ID")),
		kgo.RequestTimeoutOverhead(cfg.GetDuration("Kafka.REQUEST_TIMEOUT_OVERHEAD") * time.Second),
		kgo.ConnIdleTimeout(cfg.GetDuration("Kafka.CONNECTION_IDLE_TIMEOUT") * time.Second),
		kgo.ProducerBatchCompression(kgo.SnappyCompression()),
		kgo.ProducerBatchMaxBytes(cfg.GetInt32("Kafka.PRODUCER_BATCH_MAX_BYTES")),
		kgo.ProducerLinger(cfg.GetDuration("Kafka.PRODUCER_LINGER") * time.Millisecond),
		kgo.RequiredAcks(kgo.AllISRAcks()),
		kgo.MaxConcurrentFetches(cfg.GetInt("Kafka.MAX_CONCURRENT_FETCHES")),
	}

	if cfg.GetBool("Kafka.ALLOW_AUTO_TOPIC_CREATION") == true {
		opts = append(opts,
			kgo.AllowAutoTopicCreation(),
		)
	}

	client, err = kgo.NewClient(opts...)

	ctx, cancel := context.WithTimeout(ctx, cfg.GetDuration("Kafka.CONN_TIMEOUT")*time.Second)
	defer cancel()

	if err = client.Ping(ctx); err != nil {
		client.Close()
		return nil, err
	}

	return
}
