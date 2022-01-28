package config

import "os"

func DefaultHost() string {
	return os.Getenv("DEFAULT_HOST")
}
func RedisAddress() string {
	return os.Getenv("REDIS_ADDRESS")
}
func RedisPassword() string {
	return os.Getenv("REDIS_PASSWORD")
}
