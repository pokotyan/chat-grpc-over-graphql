package endpoint

import (
	"context"
	"grpc-server/pb"
)

func DecodeMessageRequest(_ context.Context, request *pb.MessageRequest) (*pb.MessageRequest, error) {
	return &pb.MessageRequest{Message: request.Message}, nil
}

func EncodeMessageResponse(_ context.Context, message string) (*pb.MessageResponse, error) {
	return &pb.MessageResponse{Message: message}, nil
}
