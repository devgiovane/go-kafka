package main

import (
    "github.com/confluentinc/confluent-kafka-go/kafka"
    "log"
)

func main()  {
    deliveryChan := make(chan kafka.Event)
    p := NewKafkaProducer()
    Publish("confirmed", "teste", p, []byte("transfer"), deliveryChan)
    go DeliveryReport(deliveryChan)
    p.Flush(2000)
}

func NewKafkaProducer() *kafka.Producer {
    configMap := &kafka.ConfigMap{
        "bootstrap.servers": "kafka:9092",
    }
    p, err := kafka.NewProducer(configMap)
    if err != nil {
        log.Fatalln(err)
    }
    return p
}

func Publish(message string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) {
    kMessage := kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          []byte(message),
        Key:            key,
    }
    err := producer.Produce(&kMessage, deliveryChan)
    if err != nil {
        log.Fatalln(err)
    }
}

func DeliveryReport(deliveryChan chan kafka.Event)  {
    for e := range deliveryChan {
        switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    log.Println("Error to send message")
                    return
                }
                log.Println("Success to send message", ev.TopicPartition)
        }
    }
}
