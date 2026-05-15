package service

import (
	"context"
	"fmt"
	"sync"

	maintenancev1 "go-kozadayev-exercise/task10/api/proto"
	"go-kozadayev-exercise/task10/internal/excel"
)

type MaintenanceService struct {
	maintenancev1.UnimplementedMaintenanceServiceServer
	mu     sync.RWMutex
	orders map[string]*maintenancev1.WorkOrder
}

func NewMaintenanceService() *MaintenanceService {
	return &MaintenanceService{
		orders: make(map[string]*maintenancev1.WorkOrder),
	}
}

func (s *MaintenanceService) CreateWorkOrder(ctx context.Context, req *maintenancev1.CreateWorkOrderRequest) (*maintenancev1.WorkOrder, error) {
	// Валидация входных данных
	if req.CarId == "" {
		return nil, fmt.Errorf("car_id is required")
	}
	if req.Description == "" {
		return nil, fmt.Errorf("description is required")
	}
	if req.TotalCost < 0 {
		return nil, fmt.Errorf("total_cost must be non-negative")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	id := fmt.Sprintf("WO-%d", len(s.orders)+1)
	order := &maintenancev1.WorkOrder{
		Id:          id,
		CarId:       req.CarId,
		Description: req.Description,
		TotalCost:   req.TotalCost,
	}
	s.orders[id] = order
	return order, nil
}

func (s *MaintenanceService) GetWorkOrder(ctx context.Context, req *maintenancev1.GetWorkOrderRequest) (*maintenancev1.WorkOrder, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	order, ok := s.orders[req.Id]
	if !ok {
		return nil, fmt.Errorf("work order not found")
	}
	return order, nil
}

func (s *MaintenanceService) ExportToExcel(ctx context.Context, req *maintenancev1.ExportRequest) (*maintenancev1.ExportResponse, error) {
	s.mu.RLock()
	var orders []*maintenancev1.WorkOrder
	for _, o := range s.orders {
		orders = append(orders, o)
	}
	s.mu.RUnlock()

	fileName, err := excel.ExportWorkOrders(orders)
	if err != nil {
		return nil, err
	}

	return &maintenancev1.ExportResponse{FileUrl: fileName}, nil
}
