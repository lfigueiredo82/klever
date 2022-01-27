package redis

import (
	"github.com/go-redis/redis"
)

func ReadAllCryptoCurrency(client *redis.Client) []string {
	result, _ := client.Keys("cryptos*").Result()
	return result
}

func GetCryptoCurrency(client *redis.Client, key string) string {
	return client.Get("cryptos:" + key).Val()
}
func CreateCryptoCurrency(client *redis.Client, cripto string, data interface{}) {
	client.Set("cryptos"+":"+cripto+":info", data, 0)
}

func UpdateCryptoCurrency(client *redis.Client, oldCriptoCode, newCryptoCode string, data interface{}) {
	val := client.Get("cryptos" + ":" + oldCriptoCode + ":info").Val()
	if val != "" {
		client.Set("cryptos"+":"+newCryptoCode+":info", data, 0)
		keys := []string{"cryptos" + ":" + oldCriptoCode + ":info"}
		client.Del(keys...)
	}
}

func DeleteCryptoCurrency(client *redis.Client, cripto string) {
	keys := []string{"cryptos" + ":" + cripto + ":info", "cryptos" + ":" + cripto}
	client.Del(keys...)
}

func UpvoteCrytoCurrency(client *redis.Client, cripto string) int64 {
	return client.Incr("cryptos" + ":" + cripto).Val()
}

func DownvoteCrytoCurrency(client *redis.Client, cripto string) int64 {
	return client.Decr("cryptos" + ":" + cripto).Val()
}
