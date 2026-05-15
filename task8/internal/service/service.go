package service

import (
	"context"
	"fmt"

	rentalv1 "go-kozadayev-exercise/task8/api/proto"
	"go-kozadayev-exercise/task8/internal/repository"

	"github.com/jackc/pgx/v5"
)

type RentalService struct {
	rentalv1.UnimplementedRentalServiceServer
	db   *pgx.Conn
	repo repository.Querier
}

func NewRentalService(db *pgx.Conn, repo repository.Querier) *RentalService {
	return &RentalService{db: db, repo: repo}
}

func (s *RentalService) CreateBooking(ctx context.Context, req *rentalv1.CreateBookingRequest) (*rentalv1.Booking, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := repository.New(tx)

	_, err = qtx.GetCarForUpdate(ctx, req.CarId)
	if err != nil {
		return nil, err
	}

	overlapped, err := qtx.CheckOverlap(ctx, repository.CheckOverlapParams{
		CarID:     req.CarId,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})
	if err != nil {
		return nil, err
	}
	if overlapped {
		return nil, fmt.Errorf("car is already booked for these dates")
	}

	b, err := qtx.CreateBooking(ctx, repository.CreateBookingParams{
		CarID:     req.CarId,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &rentalv1.Booking{
		Id:        b.ID,
		CarId:     b.CarID,
		StartDate: b.StartDate,
		EndDate:   b.EndDate,
	}, nil
}

func (s *RentalService) CheckAvailability(ctx context.Context, req *rentalv1.CheckAvailabilityRequest) (*rentalv1.CheckAvailabilityResponse, error) {
	overlapped, err := s.repo.CheckOverlap(ctx, repository.CheckOverlapParams{
		CarID:     req.CarId,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})
	if err != nil {
		return nil, err
	}
	return &rentalv1.CheckAvailabilityResponse{Available: !overlapped}, nil
}
