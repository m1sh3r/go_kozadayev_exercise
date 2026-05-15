package service

import (
	"context"

	carv1 "go-kozadayev-exercise/task1/api/proto"
	"go-kozadayev-exercise/task1/internal/domain"
	"go-kozadayev-exercise/task1/internal/repository"
)

type CarService struct {
	carv1.UnimplementedCarServiceServer
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) *CarService {
	return &CarService{repo: repo}
}

func (s *CarService) CreateCar(ctx context.Context, req *carv1.CreateCarRequest) (*carv1.CreateCarResponse, error) {
	car := domain.Car{
		VIN:   req.Vin,
		Brand: req.Brand,
		Model: req.Model,
		Year:  int(req.Year),
	}
	if err := s.repo.Create(ctx, car); err != nil {
		return nil, err
	}
	return &carv1.CreateCarResponse{
		Car: &carv1.Car{
			Vin:   car.VIN,
			Brand: car.Brand,
			Model: car.Model,
			Year:  int32(car.Year),
		},
	}, nil
}

func (s *CarService) GetCar(ctx context.Context, req *carv1.GetCarRequest) (*carv1.GetCarResponse, error) {
	car, err := s.repo.Get(ctx, req.Vin)
	if err != nil {
		return nil, err
	}
	return &carv1.GetCarResponse{
		Car: &carv1.Car{
			Vin:   car.VIN,
			Brand: car.Brand,
			Model: car.Model,
			Year:  int32(car.Year),
		},
	}, nil
}
