package repository

import (
	"context"
)

type Querier interface {
	CheckOverlap(ctx context.Context, arg CheckOverlapParams) (bool, error)
	CreateBooking(ctx context.Context, arg CreateBookingParams) (Booking, error)
	GetCarForUpdate(ctx context.Context, id int64) (Car, error)
}

var _ Querier = (*Queries)(nil)
