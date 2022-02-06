package endpoint

import (
	"context"
	"grpc-server/pb"
)

func DecodeMathRequest(_ context.Context, request *pb.MathRequest) (*pb.MathRequest, error) {
	return &pb.MathRequest{NumA: request.NumA, NumB: request.NumB}, nil
}

func EncodeMathResponse(_ context.Context, result float32) (*pb.MathResponse, error) {
	return &pb.MathResponse{Result: result}, nil
}
