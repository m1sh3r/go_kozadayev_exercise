package repository

import (
	"context"
)

const checkOverlap = `-- name: CheckOverlap :one
SELECT EXISTS (
    SELECT 1 FROM bookings 
    WHERE car_id = $1 AND start_date < $3 AND $2 < end_date
) as overlapped
`

type CheckOverlapParams struct {
	CarID     int64
	EndDate   int64
	StartDate int64
}

func (q *Queries) CheckOverlap(ctx context.Context, arg CheckOverlapParams) (bool, error) {
	row := q.db.QueryRow(ctx, checkOverlap, arg.CarID, arg.EndDate, arg.StartDate)
	var overlapped bool
	err := row.Scan(&overlapped)
	return overlapped, err
}

const createBooking = `-- name: CreateBooking :one
INSERT INTO bookings (car_id, start_date, end_date) VALUES ($1, $2, $3) RETURNING id, car_id, start_date, end_date
`

type CreateBookingParams struct {
	CarID     int64
	StartDate int64
	EndDate   int64
}

func (q *Queries) CreateBooking(ctx context.Context, arg CreateBookingParams) (Booking, error) {
	row := q.db.QueryRow(ctx, createBooking, arg.CarID, arg.StartDate, arg.EndDate)
	var i Booking
	err := row.Scan(
		&i.ID,
		&i.CarID,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const getCarForUpdate = `-- name: GetCarForUpdate :one
SELECT id, brand FROM cars WHERE id = $1 FOR UPDATE
`

func (q *Queries) GetCarForUpdate(ctx context.Context, id int64) (Car, error) {
	row := q.db.QueryRow(ctx, getCarForUpdate, id)
	var i Car
	err := row.Scan(&i.ID, &i.Brand)
	return i, err
}
