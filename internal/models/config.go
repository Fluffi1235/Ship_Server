package models

import "time"

type Config struct {
	ServicePort int           `yaml:"port"`
	ConnectDB   string        `yaml:"connectDB"`
	TokenTTl    time.Duration `yaml:"token_ttl"`
	TimeoutTx   int           `yaml:"timeoutTx"`
}
