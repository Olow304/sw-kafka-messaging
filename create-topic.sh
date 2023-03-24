#!/bin/bash

set -e

# Wait for Kafka to become available
until $(nc -zv kafka 9092); do
    printf 'kafka is still not available. Retrying...\n'
    sleep 5
done

# Create the "messages" topic
kafka-topics --create --bootstrap-server kafka:9092 --replication-factor 1 --partitions 1 --topic messages
