package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func InitRedis() {

	//create a new redis client
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := client.Ping(context.Background()); err.Err() != nil {
		log.Fatal("Failed to connect to Redis")
	} else {
		log.Println("connected to redis")
	}

}

func Set(ctx context.Context, key, value string) error {
	err := client.Set(ctx, key, value, 0)
	if err != nil {
		log.Printf("Error Storing in redis: %v", err)
	}
	return err.Err()
}

func Get(ctx context.Context, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Printf("Key does not exist %v,\n", key)
	} else if err != nil {
		log.Printf("Failure in Getting value from redis %v\n", err)
	}
	return val, err
}
