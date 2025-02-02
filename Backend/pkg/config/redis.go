package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

// RedisClient holds the Redis connection client
var RedisClient *redis.Client

// InitializeRedis establishes a connection to the Redis server.
func InitRedis() {
	// You can replace these values with environment variables or config options
	redisAddr := "redis:6379" // Default Redis host and port (adjust as needed)
	redisPassword := ""       // No password by default (adjust if needed)
	redisDB := 0              // Default DB

	// Initialize Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
		Protocol: 2,
	})

	// Test the Redis connection
	_, err := RedisClient.Ping(context.Background()).Result() // Use context.Background() here
	if err != nil {
		log.Fatalf("could not connect to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis")
}
