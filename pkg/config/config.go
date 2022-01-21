package config

import "os"

func RedisAddress() string {
	return os.Getenv("REDIS_ADDRESS")
}
func RedisPassword() string {
	return os.Getenv("REDIS_PASSWORD")
}
