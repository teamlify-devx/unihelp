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
		kgo.SeedBrokers(cfg.GetStringSlice("kafka.BROKERS")...),
		kgo.ClientID(cfg.GetString("kafka.CLIENT_ID")),
		kgo.RequestTimeoutOverhead(cfg.GetDuration("kafka.REQUEST_TIMEOUT_OVERHEAD") * time.Second),
		kgo.ConnIdleTimeout(cfg.GetDuration("kafka.CONNECTION_IDLE_TIMEOUT") * time.Second),
		kgo.ProducerBatchCompression(kgo.SnappyCompression()),
		kgo.ProducerBatchMaxBytes(cfg.GetInt32("kafka.PRODUCER_BATCH_MAX_BYTES")),
		kgo.ProducerLinger(cfg.GetDuration("kafka.PRODUCER_LINGER") * time.Millisecond),
		kgo.RequiredAcks(kgo.AllISRAcks()),
		kgo.MaxConcurrentFetches(cfg.GetInt("kafka.MAX_CONCURRENT_FETCHES")),
	}

	if cfg.GetBool("kafka.ALLOW_AUTO_TOPIC_CREATION") == true {
		opts = append(opts,
			kgo.AllowAutoTopicCreation(),
		)
	}

	if cfg.GetBool("kafka.ENABLE_CONSUMER") == true {
		opts = append(opts,
			kgo.ConsumerGroup(cfg.GetString("kafka.CONSUMER_GROUP")),
			kgo.ConsumeTopics(cfg.GetString("kafka.TOPICS")),
		)
	}

	if cfg.GetBool("kafka.DISABLE_AUTO_COMMIT") == true {
		opts = append(opts,
			kgo.DisableAutoCommit(),
		)
	}

	client, err = kgo.NewClient(opts...)

	ctx, cancel := context.WithTimeout(ctx, cfg.GetDuration("kafka.CONN_TIMEOUT")*time.Second)
	defer cancel()

	if err = client.Ping(ctx); err != nil {
		client.Close()
		return nil, err
	}

	return
}
