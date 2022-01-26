package redis

import (
	"github.com/go-redis/redis"
)

func ReadAllCryptoCurrency(client *redis.Client) []string {
	result, _ := client.Keys("*").Result()
	return result
}
func CreateCryptoCurrency(client *redis.Client, cripto string, data interface{}) {
	client.Set(cripto, data, 0)
}

func UpdateCryptoCurrency(client *redis.Client, oldCriptoName, newCryptoName string, value int64) {
	val := client.Get(oldCriptoName).Val()
	client.Set(newCryptoName, val, 0)
	DeleteCryptoCurrency(client, oldCriptoName)
}

func DeleteCryptoCurrency(client *redis.Client, cripto string) {
	keys := []string{cripto}
	client.Del(keys...)
}

func UpvoteCrytoCurrency(client *redis.Client, cripto string) int64 {
	return client.Incr(cripto).Val()
}

func DownvoteCrytoCurrency(client *redis.Client, cripto string) int64 {
	return client.Decr(cripto).Val()
}
