# Kafka message broker with go

### About

Concepts:

- Producer
- Consumer
- Partitions
- Keys

### Commands

```bash
go run cmd/consumer/main.go
go run cmd/producer/main.go 
```

### Commands kafka

```bash
# create topic
kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3
# list topics
kafka-topics --list --bootstrap-server=localhost:9092
# describ topic
kafka-topics --bootstrap-server=localhost:9092 --topic=teste --describe
# consumer
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste
# consumer all
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning
# consumer group
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --group=x
# producer
kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste
```

