package repository

import (
	"context"
	"fmt"
	"sync"

	"go-kozadayev-exercise/task1/internal/domain"
)

type CarRepository interface {
	Create(ctx context.Context, car domain.Car) error
	Get(ctx context.Context, vin string) (domain.Car, error)
}

type memoryRepo struct {
	mu   sync.RWMutex
	cars map[string]domain.Car
}

func NewMemoryRepository() CarRepository {
	repo := &memoryRepo{
		cars: make(map[string]domain.Car),
	}
	
	repo.cars["VIN123456789FOCUS"] = domain.Car{
		VIN:   "VIN123456789FOCUS",
		Brand: "Ford",
		Model: "Focus",
		Year:  2010,
	}
	
	repo.cars["VIN123456789LANCER"] = domain.Car{
		VIN:   "VIN123456789LANCER",
		Brand: "Mitsubishi",
		Model: "Lancer",
		Year:  2010,
	}
	
	return repo
}

func (r *memoryRepo) Create(ctx context.Context, car domain.Car) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cars[car.VIN] = car
	return nil
}

func (r *memoryRepo) Get(ctx context.Context, vin string) (domain.Car, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	car, ok := r.cars[vin]
	if !ok {
		return domain.Car{}, fmt.Errorf("car not found")
	}
	return car, nil
}
