package handler

import (
	"context"
	"grpc-server/service"

	"grpc-server/endpoint"
	"grpc-server/pb"
)

type mathHandler struct {
	s service.MathService
}

func NewMathHandler(s service.MathService) pb.MathServiceServer {
	return &mathHandler{
		s: s,
	}
}

func (h *mathHandler) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	request, err := endpoint.DecodeMathRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	res, err := h.s.Add(ctx, request.GetNumA(), request.GetNumB())
	if err != nil {
		return nil, err
	}

	response, err := endpoint.EncodeMathResponse(ctx, res)
	if err != nil {
		return nil, err
	}

	return response, nil
}
