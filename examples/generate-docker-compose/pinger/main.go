package pinger

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Main function of hello-world
func Main(client *redis.Client) (err error) {
	defer client.Close()

	if err := client.Ping().Err(); err != nil {
		fmt.Printf("Ping: %s\n", err.Error())
	} else {
		fmt.Println("Ping success")
	}
	return
}

// NewRedisClient [constructor]
func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redispass",
	})
}
