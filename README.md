# kafka-pipeline

# Pub/Sub with Kafka

This project demonstrates the implementation of a Pub/Sub system using Apache Kafka. The system consists of a producer and consumer application written in Go, along with a Kafka broker and a Zookeeper service. The setup is containerized using Docker and orchestrated with Docker Compose.

## Project Structure

The project repository includes the following files:

- **docker-compose.yml**: Defines the services required for the Kafka setup and Go program.
- **Dockerfile**: Multi-stage Dockerfile that builds the Go program and executes it in two different stages.
- **go.mod**: Go module dependencies.
- **go.sum**: Go module checksum file for verifying module consistency.
- **main.go**: The Go source code implementing the producer and consumer functionality.

## Services in `docker-compose.yml`

The `docker-compose.yml` file defines the following services:

1. **Kafka UI**: A user interface for interacting with Kafka.
2. **Kafka Broker**: The Kafka message broker.
3. **Zookeeper**: The Zookeeper service used for Kafka's coordination.
4. **Go Program**: A Go application that starts both the producer and the consumer. The producer sends messages to the Kafka topic, and the consumer reads messages from it.


