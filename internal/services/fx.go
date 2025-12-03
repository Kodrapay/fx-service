package services

import (
	"context"

	"github.com/kodra-pay/fx-service/internal/dto"
)

type FXService struct{}

func NewFXService() *FXService { return &FXService{} }

func (s *FXService) Rates(_ context.Context) []dto.FXRate {
	return []dto.FXRate{
		{Pair: "USD/NGN", Rate: 1500.0},
		{Pair: "EUR/NGN", Rate: 1620.0},
	}
}

func (s *FXService) Convert(_ context.Context, req dto.FXConvertRequest) dto.FXConvertResponse {
	rate := 1.0
	return dto.FXConvertResponse{
		From:      req.From,
		To:        req.To,
		Amount:    req.Amount,
		Rate:      rate,
		Converted: req.Amount * rate,
	}
}
