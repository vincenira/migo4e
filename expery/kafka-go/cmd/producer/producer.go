package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

	notification := models.Notification{
		From:    fromUser,
		To:      toUser,
		Message: message,
	}

	notificationJson, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: KafkaTopic,
		Key:   sarama.StringEncoder(strconv.Itoa(toUser.ID)),
		Value: sarama.StringEncoder(notificationJson),
	}
	_, _, err = producer.SendMessage(msg)
	return err
}
