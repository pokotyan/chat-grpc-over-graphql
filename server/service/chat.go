package service

import (
	"context"
	"grpc-server/pb"
	"log"
)

type chatService struct {
	messages []string
}

type ChatService interface {
	GetMessages(pb.ChatService_GetMessagesServer) error
	CreateMessage(context.Context, string) (string, error)
}

func NewChatService() ChatService {
	return &chatService{}
}

func (s *chatService) GetMessages(stream pb.ChatService_GetMessagesServer) error {
	for _, m := range s.messages {
		if err := stream.Send(&pb.MessageResponse{Message: m}); err != nil {
			return err
		}
	}
	previousCount := len(s.messages)
	for {
		currentCount := len(s.messages)
		if previousCount < currentCount && currentCount > 0 {
			m := s.messages[currentCount-1]
			log.Printf("Sent: %s", m)
			if err := stream.Send(&pb.MessageResponse{Message: m}); err != nil {
				return err
			}
		}
		previousCount = currentCount
	}
}

func (s *chatService) CreateMessage(ctx context.Context, message string) (string, error) {
	s.messages = append(s.messages, message)
	return message, nil
}
