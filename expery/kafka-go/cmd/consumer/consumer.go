package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"kafka-go/pkg/models"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

const (
	ConsumerGroup      = "notifications-group"
	ConsumerTopic      = "notifications"
	ConsumerPort       = ":8081"
	KafkaServerAddress = "localhost:9092"
)

// Helpers function
var ErrnoMessagesFound = errors.New("no messages found")

func getUserIDFromRequest(ctx *gin.Context) (string, error) {
	userID := ctx.Param("userID")
	if userID == "" {
		return "", ErrnoMessagesFound
	}
	return userID, nil
}

// Notification storage
type UserNotifications map[string][]models.Notification

type NotificationStore struct {
	data UserNotifications
	mu   sync.RWMutex
}

func (ns *NotificationStore) Add(userID string, notification models.Notification) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.data[userID] = append(ns.data[userID], notification)
}

func (ns *NotificationStore) Get(userID string) []models.Notification {
	ns.mu.RLock()
	defer ns.mu.RUnlock()
	return ns.data[userID]
}

// KAFKA RELATED FUNCTIONS

type Consumer struct {
	store *NotificationStore
}

// Setup() and Cleanup() methods are required to satisfy the sarama.ConsumerGroupHandler interface
// while they will NOT be used in this tutorial, they can serve potential rols for initialization and
// cleanup during message consumption but act as placeholders here.
func (*Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

// in the ConsumeClaim()method, The consumer listens for new messages on the topic.
// For each message, it fetches the userID (the key of the message), un-marshals the message into
// a Notification struct, and adds the notification to the NotificationStore
func (consumer *Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		userID := string(msg.Key)
		var notification models.Notification
		err := json.Unmarshal(msg.Value, &notification)
		if err != nil {
			log.Printf("failed to unmarshal notification: %v", err)
			continue
		}
		consumer.store.Add(userID, notification)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func initializeconsumerGroup() (sarama.ConsumerGroup, error) {
	// Initializes a new default configuration for Kafka
	config := sarama.NewConfig()

	// Creates a new Kafka consumer group that connects to the broker running on localhost:9092
	// the group name is "notifications-group"
	consumerGroup, err := sarama.NewConsumerGroup([]string{KafkaServerAddress}, ConsumerGroup, config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}
	return consumerGroup, nil
}

// Set up the kafka consumer group, listtens for incoming messages,and processes them using the Consumer struct methods
func setupConsumerGroup(ctx context.Context, store *NotificationStore) {
	consumerGroup, err := initializeconsumerGroup()
	if err != nil {
		log.Printf("Initialization error: %v", err)
	}
	defer consumerGroup.Close()

	consumer := &Consumer{
		store: store,
	}

	for {
		err = consumerGroup.Consume(ctx, []string{ConsumerTopic}, consumer)
		if err != nil {
			log.Printf("error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}

// Fetches the notifications for the provide userId from the NotificationSore.
func handleNotifications(ctx *gin.Context, store *NotificationStore) {
	userID, err := getUserIDFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	notes := store.Get(userID)
	if len(notes) == 0 {
		ctx.JSON(http.StatusOK,
			gin.H{
				"message":       "No notification found for user",
				"notifications": []models.Notification{},
			})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"notifications": notes})
}

func main() {
	// Creates an instance of NotificationStore to hold the notifications
	store := &NotificationStore{
		data: make(UserNotifications),
	}

	// Sets up a cancellable context that can be used to stop the consumer group
	ctx, cancel := context.WithCancel(context.Background())

	// Starts the consumer group in a separate Goroutine, allowing it to operate concurrently
	// without blocking the main thread
	go setupConsumerGroup(ctx, store)
	defer cancel()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// Define a GET endpoint /notifications/:userID that will fetch the notifications for a specific user via the handleNotifications() function when accessed
	router.GET("/notifications/:userID", func(ctx *gin.Context) {
		handleNotifications(ctx, store)
	})

	fmt.Printf("Kafka CONSUMER (Group: %s) 👥📥"+"started at http://localhost%s\n", ConsumerGroup, ConsumerPort)
	if err := router.Run(ConsumerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
