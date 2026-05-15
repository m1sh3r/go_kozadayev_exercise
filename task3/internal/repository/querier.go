package repository

import (
	"context"
)

type Querier interface {
	AssembleCar(ctx context.Context, arg AssembleCarParams) (Car, error)
	GetCarSpec(ctx context.Context, id int64) (GetCarSpecRow, error)
}

var _ Querier = (*Queries)(nil)
