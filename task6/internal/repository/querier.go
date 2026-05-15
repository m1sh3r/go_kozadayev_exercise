package repository

import (
	"context"
)

type Querier interface {
	CreateCar(ctx context.Context, arg CreateCarParams) (Car, error)
	CreateOwner(ctx context.Context, name string) (Owner, error)
	CreateServiceRecord(ctx context.Context, arg CreateServiceRecordParams) (ServiceRecord, error)
	GetDashboardData(ctx context.Context, ownerID int64) ([]GetDashboardDataRow, error)
	GetOwner(ctx context.Context, id int64) (Owner, error)
}

var _ Querier = (*Queries)(nil)
