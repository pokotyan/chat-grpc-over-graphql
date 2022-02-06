package service

import (
	"context"
)

type mathService struct{}

type MathService interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
}

func NewMathService() MathService {
	return &mathService{}
}

func (s *mathService) Add(ctx context.Context, numA, numB float32) (float32, error) {
	return numA + numB, nil
}
