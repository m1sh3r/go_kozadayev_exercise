package service

import (
	"context"
	"time"

	dashboardv1 "go-kozadayev-exercise/task6/api/proto"
	"go-kozadayev-exercise/task6/internal/repository"
)

type DashboardService struct {
	dashboardv1.UnimplementedDashboardServiceServer
	repo repository.Querier
}

func NewDashboardService(repo repository.Querier) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetDashboard(ctx context.Context, req *dashboardv1.GetDashboardRequest) (*dashboardv1.GetDashboardResponse, error) {
	owner, err := s.repo.GetOwner(ctx, req.OwnerId)
	if err != nil {
		return nil, err
	}

	rows, err := s.repo.GetDashboardData(ctx, req.OwnerId)
	if err != nil {
		return nil, err
	}

	entries := make([]*dashboardv1.CarDashboardEntry, len(rows))
	for i, row := range rows {
		entry := &dashboardv1.CarDashboardEntry{
			Car: &dashboardv1.Car{
				Id:      row.CarID,
				Brand:   row.CarBrand,
				OwnerId: owner.ID,
			},
		}
		if row.ServiceID != 0 {
			entry.LastService = &dashboardv1.ServiceRecord{
				Id:          row.ServiceID,
				CarId:       row.CarID,
				Description: row.ServiceDesc,
				Date:        row.ServiceDate,
			}
		}
		entries[i] = entry
	}

	return &dashboardv1.GetDashboardResponse{
		Owner: &dashboardv1.Owner{Id: owner.ID, Name: owner.Name},
		Cars:  entries,
	}, nil
}

func (s *DashboardService) CreateOwner(ctx context.Context, req *dashboardv1.CreateOwnerRequest) (*dashboardv1.Owner, error) {
	o, err := s.repo.CreateOwner(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &dashboardv1.Owner{Id: o.ID, Name: o.Name}, nil
}

func (s *DashboardService) CreateCar(ctx context.Context, req *dashboardv1.CreateCarRequest) (*dashboardv1.Car, error) {
	c, err := s.repo.CreateCar(ctx, repository.CreateCarParams{Brand: req.Brand, OwnerID: req.OwnerId})
	if err != nil {
		return nil, err
	}
	return &dashboardv1.Car{Id: c.ID, Brand: c.Brand, OwnerId: c.OwnerID}, nil
}

func (s *DashboardService) CreateServiceRecord(ctx context.Context, req *dashboardv1.CreateServiceRecordRequest) (*dashboardv1.ServiceRecord, error) {
	sr, err := s.repo.CreateServiceRecord(ctx, repository.CreateServiceRecordParams{
		CarID:       req.CarId,
		Description: req.Description,
		Date:        time.Now().Unix(),
	})
	if err != nil {
		return nil, err
	}
	return &dashboardv1.ServiceRecord{Id: sr.ID, CarId: sr.CarID, Description: sr.Description, Date: sr.Date}, nil
}
