package franz

import (
	"context"
	"time"

	cfg "github.com/spf13/viper"
	"github.com/twmb/franz-go/pkg/kgo"
)

// NewKafkaClient Return new Kafka client instance
/*
Example Config yaml:

	kafka:
  BROKERS:
    - "localhost:19092"
    - "localhost:19093"
    - "localhost:19094"
  TOPICS:
    - "audit-events"
  CLIENT_ID: "audit-log-client"
  REQUEST_TIMEOUT_OVERHEAD: 10
  CONNECTION_IDLE_TIMEOUT: 10
  PRODUCER_BATCH_MAX_BYTES: 1000000
  PRODUCER_LINGER: 10
  MAX_CONCURRENT_FETCHES: 3
  ALLOW_AUTO_TOPIC_CREATION: true
  CONN_TIMEOUT: 10
  DISABLE_AUTO_COMMIT: true
  ENABLE_PRODUCER: false
  ENABLE_CONSUMER: true
  CONSUMER_GROUP: "audit-consumer-group"
  ENABLE_TUNING: true
  TUNING:
    FETCH_MIN_BYTES: 1e5
    FETCH_MAX_WAIT: 50
    MAX_CONCURRENT_FETCHES: 10
*/
func NewKafkaClient(ctx context.Context) (client *kgo.Client, err error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(cfg.GetStringSlice("kafka.BROKERS")...),
		kgo.ClientID(cfg.GetString("kafka.CLIENT_ID")),
		kgo.RequestTimeoutOverhead(cfg.GetDuration("kafka.REQUEST_TIMEOUT_OVERHEAD") * time.Second),
		kgo.ConnIdleTimeout(cfg.GetDuration("kafka.CONNECTION_IDLE_TIMEOUT") * time.Second),
	}

	if cfg.GetBool("kafka.ALLOW_AUTO_TOPIC_CREATION") == true {
		opts = append(opts,
			kgo.AllowAutoTopicCreation(),
		)
	}

	if cfg.GetBool("kafka.ENABLE_PRODUCER") == true {
		opts = append(opts,
			kgo.ProducerBatchCompression(kgo.SnappyCompression()),
			kgo.ProducerBatchMaxBytes(cfg.GetInt32("kafka.PRODUCER_BATCH_MAX_BYTES")),
			kgo.ProducerLinger(cfg.GetDuration("kafka.PRODUCER_LINGER")*time.Millisecond),
			kgo.RequiredAcks(kgo.AllISRAcks()),
		)
	}

	if cfg.GetBool("kafka.ENABLE_CONSUMER") == true {
		opts = append(opts,
			kgo.ConsumerGroup(cfg.GetString("kafka.CONSUMER_GROUP")),
			kgo.ConsumeTopics(cfg.GetString("kafka.TOPICS")),
		)
	}

	if cfg.GetBool("kafka.ENABLE_TUNING") == true {
		opts = append(opts,
			kgo.FetchMinBytes(cfg.GetInt32("kafka.tuning.FETCH_MIN_BYTES")),
			kgo.FetchMaxWait(cfg.GetDuration("kafka.tuning.FETCH_MAX_WAIT")*time.Millisecond),
			kgo.MaxConcurrentFetches(cfg.GetInt("kafka.tuning.MAX_CONCURRENT_FETCHES")),
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
