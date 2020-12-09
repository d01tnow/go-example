package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func fatalOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
func getValue(ctx context.Context, rdb *redis.Client, key string) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("%s does not exist\n", key)
	} else if err != nil {
		fatalOnError(err, "get value failed")
	} else {
		fmt.Printf("{%s: %s}\n", key, val)
	}
}

func basicClient() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.50.38:6379",
		Password: "",
		DB:       0,
	})
	status := rdb.Ping(ctx)
	fatalOnError(status.Err(), "Ping failed")
	err := rdb.Set(ctx, "k1", "v1", 10*time.Second).Err()
	fatalOnError(err, "Set failed")
	getValue(ctx, rdb, "k1")
	getValue(ctx, rdb, "k2")

}

func main() {
	basicClient()

}
