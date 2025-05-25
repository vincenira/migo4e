# tutorial To test the real-time notification system

1. Download the bitnami's Kafka

```sh
mkdir toto-kafka && cd toto-kafka
curl -sSL \
https://raw.githubusercontent.com/bitnami/containers/main/bitnami/kafka/docker-compose.yml > docker-compose.yml

```

2. Change the below in the docker-compose file

`KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092`

3. start up the docker-compose
   `docker compose up -d`

4. Start the producer
   `go run cmd/producer/producer.go`
5. start the consumer
   `go run cmd/consumer/consumer.go`
6. Sending notifications

- user 1(Emma) receives a notification from User2(Bruno)
  `curl -X POST http://localhost:8080/send \
-d "fromID=2&toID=1&message=Bruno started following you."
`
- user 2(Bruno) receives a notification from User 1(Emma):
  `curl -X POST http://localhost:8080/send \
-d "fromID=1&toID=2&message=Emma mentioned you in a comment: 'Great seeing you yesterday, @Bruno!'"
`
- User 1(Emma) receives a notification from User4 (lena):
  `curl -X POST http://localhost:8080/send \
-d "fromID=4&toID=1&message=Lena liked your post: 'My weekend getaway!'"
`
- Retrieving notifications
  `curl http://localhost:8081/notifcations/1`
