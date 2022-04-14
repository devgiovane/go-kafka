package main

import (
    "github.com/confluentinc/confluent-kafka-go/kafka"
    "log"
)

func main()  {
    configMap := &kafka.ConfigMap{
        "bootstrap.servers": "kafka:9092",
        "client.id": "app-consumer",
        "group.id": "app-group",
    }
    c, err := kafka.NewConsumer(configMap)
    if err != nil {
        log.Fatalln(err)
    }
    topics := []string{"teste"}
    c.SubscribeTopics(topics, nil)
    for {
        msg, err := c.ReadMessage(-1)
        if err != nil {
            log.Fatalln(err)
        }
        log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
    }
}


