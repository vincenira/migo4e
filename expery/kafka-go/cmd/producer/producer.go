package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"kafka-go/pkg/models"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

const (
	ProducerPort       = ":8080"
	KafkaServerAddress = "localhost:9092"
	KafkaTopic         = "notifications"
)

// Helpers
var ErrUserNotFoundInProducer = errors.New("user not found")

func findUserById(id int, users []models.User) (models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, ErrUserNotFoundInProducer
}

func getIdFromRequest(formValue string, ctx *gin.Context) (int, error) {
	id, err := strconv.Atoi(ctx.PostForm(formValue))
	if err != nil {
		return 0, fmt.Errorf("Failed to parse ID from form value %s: %w", formValue, err)
	}
	return id, nil
}

// KAFKA FUNCTIONS

func sendKafkaMessage(producer sarama.SyncProducer, users []models.User, ctx *gin.Context, fromID, toID int) error {
	message := ctx.PostForm("message")
	fromUser, err := findUserById(fromID, users)
	if err != nil {
		return err
	}

	toUser, err := findUserById(toID, users)
	if err != nil {
		return err
	}
	// initializes a notification struct that encapsulates information about the sender, the recipient,
	// and the actual message.
	notification := models.Notification{
		From:    fromUser,
		To:      toUser,
		Message: message,
	}

	notificationJson, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	// Constrcuts a ProducerMessage for the notifications topic, setting the recipient's ID as the key
	// and the message content, which is the serialized form of the Notification as the value
	msg := &sarama.ProducerMessage{
		Topic: KafkaTopic,
		Key:   sarama.StringEncoder(strconv.Itoa(toUser.ID)),
		Value: sarama.StringEncoder(notificationJson),
	}
	// Sends the constructed message to the "notifications" topic
	_, _, err = producer.SendMessage(msg)
	return err
}

func sendMessageHandler(producer sarama.SyncProducer, users []models.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fromID, err := getIdFromRequest("fromID", ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		toID, err := getIdFromRequest("toID", ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err = sendKafkaMessage(producer, users, ctx, fromID, toID)
		if errors.Is(err, ErrUserNotFoundInProducer) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Notification sent successfully!"})
	}
}

func setupProducer() (sarama.SyncProducer, error) {
	// Initializes a new default configuration for kafka.
	// think of it as setting up the parameters befroe connecting to the broker
	config := sarama.NewConfig()
	// Ensures that the producer receives an acknowledgment once the message is successfully
	// stored in the "notifications"
	config.Producer.Return.Successes = true
	// Initializes a synchronos kafka producer that connects to the kafka broker running at localhost:9092
	producer, err := sarama.NewSyncProducer([]string{KafkaServerAddress}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to setup producer: %w", err)
	}
	return producer, nil
}

func main() {
	users := []models.User{
		{ID: 1, Name: "Emma"},
		{ID: 2, Name: "Bruno"},
		{ID: 3, Name: "Rick"},
		{ID: 4, Name: "Lena"},
	}

	producer, err := setupProducer()
	if err != nil {
		log.Fatalf("failed to initialize producer: %v", err)
	}
	defer producer.Close()
	router := gin.Default()
	router.POST("/send", sendMessageHandler(producer, users))

	fmt.Printf("Kafka Producer started at the http://localhost%s\n", ProducerPort)

	if err := router.Run(ProducerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
