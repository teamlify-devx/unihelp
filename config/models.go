package config

import "time"

// Config of application
type Config struct {
	Server      Server      `mapstructure:"server,omitempty"`
	Logger      Logger      `mapstructure:"logger,omitempty"`
	Postgresql  Postgresql  `mapstructure:"postgresql,omitempty"`
	Redis       Redis       `mapstructure:"redis,omitempty"`
	Clickhouse  Clickhouse  `mapstructure:"clickhouse,omitempty"`
	Queue       Queue       `mapstructure:"queue,omitempty"`
	PubSub      PubSub      `mapstructure:"pubsub,omitempty"`
	Nats        Nats        `mapstructure:"nats,omitempty"`
	MinIoBucket MinIoBucket `mapstructure:"minio-bucket,omitempty"`
	Kafka       Kafka       `mapstructure:"kafka,omitempty"`
}

// Server config
type Server struct {
	SERVER_HEADER                  string        `mapstructure:"SERVER_HEADER,omitempty"`
	PROJECT_NAME                   string        `mapstructure:"PROJECT_NAME,omitempty"`
	SERVICE_NAME                   string        `mapstructure:"SERVICE_NAME,omitempty"`
	API_VER                        string        `mapstructure:"API_VER,omitempty"`
	APP_ENV                        string        `mapstructure:"APP_ENV,omitempty"`
	APP_DEBUG                      bool          `mapstructure:"APP_DEBUG,omitempty"`
	TIMEOUT                        int           `mapstructure:"TIMEOUT,omitempty"`
	APP_SECRET                     string        `mapstructure:"APP_SECRET,omitempty"`
	JWT_TOKEN_EXPIRE_TIME          int           `mapstructure:"JWT_TOKEN_EXPIRE_TIME,omitempty"`
	JWT_TOKEN_REMEMBER_EXPIRE_TIME int           `mapstructure:"JWT_TOKEN_REMEMBER_EXPIRE_TIME,omitempty"`
	APP_VERSION                    string        `mapstructure:"APP_VERSION,omitempty"`
	READ_TIMEOUT                   time.Duration `mapstructure:"READ_TIMEOUT,omitempty"`
	WRITE_TIMEOUT                  time.Duration `mapstructure:"WRITE_TIMEOUT,omitempty"`
	MAX_CONN_IDLE                  time.Duration `mapstructure:"MAX_CONN_IDLE,omitempty"`
	MAX_CONN_AGE                   time.Duration `mapstructure:"MAX_CONN_AGE,omitempty"`
	ENABLE_PROFILER                bool          `mapstructure:"ENABLE_PROFILER,omitempty"`
	ENABLE_METRICS                 bool          `mapstructure:"ENABLE_METRICS,omitempty"`
	ENABLE_LOGGER                  bool          `mapstructure:"ENABLE_LOGGER,omitempty"`
	ENABLE_DOCS                    bool          `mapstructure:"ENABLE_DOCS,omitempty"`
	ENABLE_RATE_LIMIT              bool          `mapstructure:"ENABLE_RATE_LIMIT,omitempty"`
	RATE_LIMIT_MAX                 int           `RATE_LIMIT_MAX:"MAX_CONN_AGE,omitempty"`
	RATE_LIMIT_EXP                 int           `mapstructure:"RATE_LIMIT_EXP,omitempty"`
	AUTH_API_URL                   string        `mapstructure:"AUTH_API_URL,omitempty"`
	CORS_ALLOWED_ORIGINS           []string      `mapstructure:"CORS_ALLOWED_ORIGINS,omitempty"`
}

// Http config
type Http struct {
	PORT                 string        `mapstructure:"PORT,omitempty"`
	TIMEOUT              time.Duration `mapstructure:"TIMEOUT,omitempty"`
	READ_TIMEOUT         time.Duration `mapstructure:"READ_TIMEOUT,omitempty"`
	WRITE_TIMEOUT        time.Duration `mapstructure:"WRITE_TIMEOUT,omitempty"`
	COOKIE_LIFE_TIME     int           `mapstructure:"COOKIE_LIFE_TIME,omitempty"`
	SESSION_COOKIE_NAME  string        `mapstructure:"SESSION_COOKIE_NAME,omitempty"`
	SSL_CERT_PATH        string        `mapstructure:"SSL_CERT_PATH,omitempty"`
	SSL_CERT_KEY         string        `mapstructure:"SSL_CERT_KEY,omitempty"`
	CORS_ALLOWED_ORIGINS []string      `mapstructure:"CORS_ALLOWED_ORIGINS,omitempty"`
	RATE_LIMIT_MAX       int           `mapstructure:"RATE_LIMIT_MAX,omitempty"`
	RATE_LIMIT_EXPIRY    time.Duration `mapstructure:"RATE_LIMIT_EXPIRY,omitempty"`
}

// Grpc config
type Grpc struct {
	PORT                            string        `mapstructure:"PORT,omitempty"`
	GRPC_TIMEOUT                    time.Duration `mapstructure:"GRPC_TIMEOUT,omitempty"`
	MAX_MSG_SIZE                    int           `mapstructure:"MAX_MSG_SIZE,omitempty"`
	MAX_CONN_IDLE                   time.Duration `mapstructure:"MAX_CONN_IDLE,omitempty"`
	MAX_CONN_AGE                    time.Duration `mapstructure:"MAX_CONN_AGE,omitempty"`
	ENABLE_KEEPALIVE                bool          `mapstructure:"ENABLE_KEEPALIVE,omitempty"`
	KEEPALIVE_TIME                  time.Duration `mapstructure:"KEEPALIVE_TIME,omitempty"`
	KEEPALIVE_TIMEOUT               time.Duration `mapstructure:"KEEPALIVE_TIMEOUT,omitempty"`
	KEEPALIVE_PERMIT_WITHOUT_STREAM bool          `mapstructure:"KEEPALIVE_PERMIT_WITHOUT_STREAM,omitempty"`
	ENABLE_TLS                      bool          `mapstructure:"ENABLE_TLS,omitempty"`
	TLS_CERT_FILE                   string        `mapstructure:"TLS_CERT_FILE,omitempty"`
	TLS_KEY_FILE                    string        `mapstructure:"TLS_KEY_FILE,omitempty"`
}

// Logger config
type Logger struct {
	DISABLE_CALLER     bool   `mapstructure:"DISABLE_CALLER,omitempty"`
	DISABLE_STACKTRACE bool   `mapstructure:"DISABLE_STACKTRACE,omitempty"`
	ENCODING           string `mapstructure:"ENCODING,omitempty"`
	LEVEL              string `mapstructure:"LEVEL,omitempty"`
	LOG_ENV            string `mapstructure:"LOG_ENV,omitempty"`
}

// Postgresql config
type Postgresql struct {
	HOST           string `mapstructure:"HOST,omitempty"`
	PORT           int    `mapstructure:"PORT,omitempty"`
	USER           string `mapstructure:"USER,omitempty"`
	PASS           string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB     string `mapstructure:"DEFAULT_DB,omitempty"`
	DEFAULT_SCHEMA string `mapstructure:"DEFAULT_SCHEMA,omitempty"`
	MAX_CONN       int    `mapstructure:"MAX_CONN,omitempty"`
	DRIVER         string `mapstructure:"DRIVER,omitempty"`
}

// Redis config
type Redis struct {
	HOST          string `mapstructure:"HOST,omitempty"`
	PORT          int    `mapstructure:"PORT,omitempty"`
	USER          string `mapstructure:"USER,omitempty"`
	PASS          string `mapstructure:"PASS,omitempty"`
	DEFAULT_DB    int    `mapstructure:"DEFAULT_DB,omitempty"`
	MIN_IDLE_CONN int    `mapstructure:"MIN_IDLE_CONN,omitempty"`
	POOL_SIZE     int    `mapstructure:"POOL_SIZE,omitempty"`
	POOL_TIMEOUT  int    `mapstructure:"POOL_TIMEOUT,omitempty"`
	DEF_TTL       int    `mapstructure:"DEF_TTL,omitempty"`
}

// Clickhouse config
type Clickhouse struct {
	HOST          string        `mapstructure:"HOST,omitempty"`
	PORT          int           `mapstructure:"PORT,omitempty"`
	USER          string        `mapstructure:"USER,omitempty"`
	PASS          string        `mapstructure:"PASS,omitempty"`
	DEFAULT_DB    string        `mapstructure:"DEFAULT_DB,omitempty"`
	DEBUG         bool          `mapstructure:"Debug,omitempty"`
	DIAL_TIMEOUT  time.Duration `mapstructure:"DIAL_TIMEOUT,omitempty"`
	MAX_OPEN_CONN int           `mapstructure:"MAX_OPEN_CONN,omitempty"`
	MIN_IDLE_CONN int           `mapstructure:"MIN_IDLE_CONN,omitempty"`
}

// Queue config
type Queue struct {
	SMS_QUEUE_NORMAL           string `mapstructure:"SMS_QUEUE_NORMAL,omitempty"`
	SMS_QUEUE_NORMAL_TEMPLATE  string `mapstructure:"SMS_QUEUE_NORMAL_TEMPLATE,omitempty"`
	SMS_QUEUE_OTP              string `mapstructure:"SMS_QUEUE_OTP,omitempty"`
	SMS_QUEUE_OTP_TEMPLATE     string `mapstructure:"SMS_QUEUE_OTP_TEMPLATE,omitempty"`
	PUSH_QUEUE_DATA            string `mapstructure:"PUSH_QUEUE_DATA,omitempty"`
	PUSH_QUEUE_DATA_TEMPLATE   string `mapstructure:"PUSH_QUEUE_DATA_TEMPLATE,omitempty"`
	PUSH_QUEUE_NOTIFY          string `mapstructure:"PUSH_QUEUE_NOTIFY,omitempty"`
	PUSH_QUEUE_NOTIFY_TEMPLATE string `mapstructure:"PUSH_QUEUE_NOTIFY_TEMPLATE,omitempty"`
	MAIL_QUEUE_NORMAL          string `mapstructure:"MAIL_QUEUE_NORMAL,omitempty"`
	MAIL_QUEUE_NORMAL_TEMPLATE string `mapstructure:"MAIL_QUEUE_NORMAL_TEMPLATE,omitempty"`
	MAIL_QUEUE_OTP             string `mapstructure:"MAIL_QUEUE_OTP,omitempty"`
	MAIL_QUEUE_OTP_TEMPLATE    string `mapstructure:"MAIL_QUEUE_OTP_TEMPLATE,omitempty"`
}

// PubSub config
type PubSub struct {
	PROJECT_ID  string `mapstructure:"PROJECT_ID,omitempty"`
	CONFIG_PATH string `mapstructure:"CONFIG_PATH,omitempty"`
}

// Nats queue system configs
type Nats struct {
	HOST         string `json:"SERVER_HOST,omitempty"`
	CLIENT_PORT  int    `json:"CLIENT_PORT,omitempty"`
	CLUSTER_PORT int    `json:"SERVER_PORT,omitempty"`
	USER         string `json:"USER,omitempty"`
	PASS         string `json:"PASS,omitempty"`
}

// MinIoBucket config
type MinIoBucket struct {
	S3_ENDPOINT   string `mapstructure:"S3_ENDPOINT,omitempty"`
	S3_ACCESS_KEY string `mapstructure:"S3_ACCESS_KEY,omitempty"`
	S3_SECRET_KEY string `mapstructure:"S3_SECRET_KEY,omitempty"`
	S3_REGION     string `mapstructure:"S3_REGION,omitempty"`
	S3_USE_SSL    bool   `mapstructure:"S3_USE_SSL,omitempty"`
}

// Kafka config
type Kafka struct {
	BROKERS  []string `mapstructure:"BROKERS,omitempty"`
	GROUP_ID string   `mapstructure:"GROUP_ID,omitempty"`
	TOPIC    string   `mapstructure:"TOPIC,omitempty"`
}
