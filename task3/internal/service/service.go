package service

import (
	"context"

	assemblyv1 "go-kozadayev-exercise/task3/api/proto"
	"go-kozadayev-exercise/task3/internal/repository"

	"github.com/jackc/pgx/v5/pgtype"
)

type AssemblyService struct {
	assemblyv1.UnimplementedAssemblyServiceServer
	repo repository.Querier
}

func NewAssemblyService(repo repository.Querier) *AssemblyService {
	return &AssemblyService{repo: repo}
}

func (s *AssemblyService) AssembleCar(ctx context.Context, req *assemblyv1.AssembleCarRequest) (*assemblyv1.AssembleCarResponse, error) {
	_, err := s.repo.AssembleCar(ctx, repository.AssembleCarParams{
		ID:             req.CarId,
		EngineID:       pgtype.Int8{Int64: req.EngineId, Valid: true},
		TransmissionID: pgtype.Int8{Int64: req.TransmissionId, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	specRes, err := s.GetCarSpec(ctx, &assemblyv1.GetCarSpecRequest{CarId: req.CarId})
	if err != nil {
		return nil, err
	}
	return &assemblyv1.AssembleCarResponse{CarSpec: specRes.CarSpec}, nil
}

func (s *AssemblyService) GetCarSpec(ctx context.Context, req *assemblyv1.GetCarSpecRequest) (*assemblyv1.GetCarSpecResponse, error) {
	row, err := s.repo.GetCarSpec(ctx, req.CarId)
	if err != nil {
		return nil, err
	}

	spec := &assemblyv1.CarSpec{
		Id:    row.ID,
		Brand: row.Brand,
		Model: row.Model,
		Year:  row.Year,
	}

	if row.EngineID.Valid {
		spec.Engine = &assemblyv1.Engine{
			Id:     row.EngineID.Int64,
			Name:   row.EngineName.String,
			Power:  row.EnginePower.Int32,
			Volume: float32(row.EngineVolume.Float32),
		}
	}

	if row.TransmissionID.Valid {
		spec.Transmission = &assemblyv1.Transmission{
			Id:   row.TransmissionID.Int64,
			Type: row.TransmissionType.String,
		}
	}

	return &assemblyv1.GetCarSpecResponse{CarSpec: spec}, nil
}
