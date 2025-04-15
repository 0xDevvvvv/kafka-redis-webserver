package main

import (
	"context"
	"log"
	"net/http"

	"github.com/0xDevvvvv/kafka-redis-webserver/kafka"
	"github.com/0xDevvvvv/kafka-redis-webserver/redis"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

const (
	topic   = "messages"
	groupID = "group-1"
)

var brokers = []string{"localhost:9092"}

// handler for post request
func handlePush(c *gin.Context) {
	var msg Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	kafka.ProduceMessage(topic, msg.Key, msg.Value)
	c.JSON(http.StatusAccepted, gin.H{"success": "Key and value entered"})

}

// handler for get request
func handlePull(c *gin.Context) {
	key := c.Param("key")
	value, err := redis.Get(context.Background(), key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "key not found"})
		return
	}
	kafka.ProduceMessage(topic, key, value) // goroutine-3
	c.JSON(http.StatusOK, gin.H{
		"key":   key,
		"value": value,
	})

}

func main() {

	//initialize kafka and start the consumer
	kafka.InitProducer(brokers)
	kafka.StartConsumer(brokers, topic, groupID)

	//initialize redis
	redis.InitRedis()

	//set up HTTP endpoint
	router := gin.Default()
	router.POST("/push", handlePush)
	router.GET("/pull/:key", handlePull)
	log.Println("✅ Server is running at http://localhost:8080")
	router.Run(":8080")
}
