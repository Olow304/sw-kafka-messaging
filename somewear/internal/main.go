package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"net"
	"sync"

	pb "github.com/olow304/somewear/proto"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

type messagingServer struct {
	pb.UnimplementedMessagingServer
	messages    []string
	subscribers []*subscriber
	mu          sync.Mutex
	kafkaWriter *kafka.Writer
	newMessages chan string
}

type subscriber struct {
	stream pb.Messaging_StreamMessagesServer
}

func newMessagingServer(kafkaWriter *kafka.Writer) *messagingServer {
	return &messagingServer{
		messages:    []string{},
		subscribers: []*subscriber{},
		kafkaWriter: kafkaWriter,
		newMessages: make(chan string),
	}
}

func (s *messagingServer) consumeKafkaMessages() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_0_0_0

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	partitionConsumer, err := consumer.ConsumePartition("messages", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to create Kafka partition consumer: %v", err)
	}

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			s.newMessages <- string(msg.Value)
			fmt.Println("Received message from Kafka", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			log.Printf("Kafka partition consumer error: %v", err)
		}
	}
}

func (s *messagingServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	err := s.kafkaWriter.WriteMessages(ctx, kafka.Message{Value: []byte(req.Message)})
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{Status: "Message sent"}, nil
}

func (s *messagingServer) StreamMessages(stream pb.Messaging_StreamMessagesServer) error {
	sub := &subscriber{stream: stream}
	s.mu.Lock()
	s.subscribers = append(s.subscribers, sub)
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		for i, sub := range s.subscribers {
			if sub.stream == stream {
				s.subscribers = append(s.subscribers[:i], s.subscribers[i+1:]...)
				break
			}
		}
		s.mu.Unlock()
	}()

	for _, message := range s.messages {
		stream.Send(&pb.StreamMessagesResponse{Message: message})
		fmt.Println("Sent buffered message to subscriber", message)
	}

	for {
		select {
		case msg := <-s.newMessages:
			s.mu.Lock()
			s.messages = append(s.messages, msg)
			s.mu.Unlock()
			stream.Send(&pb.StreamMessagesResponse{Message: msg})
			fmt.Println("Sent message to subscriber", msg)
		case <-stream.Context().Done():
			return nil
		}
	}
}

func main() {
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "messages",
		Balancer: &kafka.LeastBytes{},
	})

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		GroupID:  "message_consumers",
		Topic:    "messages",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer kafkaWriter.Close()
	defer kafkaReader.Close()

	// Define the gRPC server struct
	server := newMessagingServer(kafkaWriter)
	go server.consumeKafkaMessages()

	// Start the gRPC server
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessagingServer(grpcServer, server)

	fmt.Println("gRPC server started on :8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
