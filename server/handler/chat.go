package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"grpc-server/service"

	"grpc-server/endpoint"
	"grpc-server/pb"
)

type chatHandler struct {
	s service.ChatService
}

func NewChatHandler(s service.ChatService) pb.ChatServiceServer {
	return &chatHandler{
		s: s,
	}
}

func (h *chatHandler) GetMessages(_ *empty.Empty, stream pb.ChatService_GetMessagesServer) error {
	if err := h.s.GetMessages(stream); err != nil {
		return err
	}
	return nil
}

func (h *chatHandler) CreateMessage(ctx context.Context, request *pb.MessageRequest) (*pb.MessageResponse, error) {
	req, err := endpoint.DecodeMessageRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	message, err := h.s.CreateMessage(ctx, req.GetMessage())
	if err != nil {
		return nil, err
	}

	response, err := endpoint.EncodeMessageResponse(ctx, message)
	if err != nil {
		return nil, err
	}
	return response, nil
}
