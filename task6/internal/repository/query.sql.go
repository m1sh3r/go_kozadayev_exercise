package repository

import (
	"context"
)

const createCar = `-- name: CreateCar :one
INSERT INTO cars (brand, owner_id) VALUES ($1, $2) RETURNING id, brand, owner_id
`

type CreateCarParams struct {
	Brand   string
	OwnerID int64
}

func (q *Queries) CreateCar(ctx context.Context, arg CreateCarParams) (Car, error) {
	row := q.db.QueryRow(ctx, createCar, arg.Brand, arg.OwnerID)
	var i Car
	err := row.Scan(&i.ID, &i.Brand, &i.OwnerID)
	return i, err
}

const createOwner = `-- name: CreateOwner :one
INSERT INTO owners (name) VALUES ($1) RETURNING id, name
`

func (q *Queries) CreateOwner(ctx context.Context, name string) (Owner, error) {
	row := q.db.QueryRow(ctx, createOwner, name)
	var i Owner
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createServiceRecord = `-- name: CreateServiceRecord :one
INSERT INTO service_records (car_id, description, date) VALUES ($1, $2, $3) RETURNING id, car_id, description, date
`

type CreateServiceRecordParams struct {
	CarID       int64
	Description string
	Date        int64
}

func (q *Queries) CreateServiceRecord(ctx context.Context, arg CreateServiceRecordParams) (ServiceRecord, error) {
	row := q.db.QueryRow(ctx, createServiceRecord, arg.CarID, arg.Description, arg.Date)
	var i ServiceRecord
	err := row.Scan(
		&i.ID,
		&i.CarID,
		&i.Description,
		&i.Date,
	)
	return i, err
}

const getDashboardData = `-- name: GetDashboardData :many
SELECT 
    c.id as car_id, c.brand as car_brand,
    sr.id as service_id, sr.description as service_desc, sr.date as service_date
FROM cars c
LEFT JOIN (
    SELECT DISTINCT ON (car_id) id, car_id, description, date
    FROM service_records
    ORDER BY car_id, date DESC
) sr ON c.id = sr.car_id
WHERE c.owner_id = $1
`

type GetDashboardDataRow struct {
	CarID       int64
	CarBrand    string
	ServiceID   int64
	ServiceDesc string
	ServiceDate int64
}

func (q *Queries) GetDashboardData(ctx context.Context, ownerID int64) ([]GetDashboardDataRow, error) {
	rows, err := q.db.Query(ctx, getDashboardData, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDashboardDataRow
	for rows.Next() {
		var i GetDashboardDataRow
		if err := rows.Scan(
			&i.CarID,
			&i.CarBrand,
			&i.ServiceID,
			&i.ServiceDesc,
			&i.ServiceDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOwner = `-- name: GetOwner :one
SELECT id, name FROM owners WHERE id = $1
`

func (q *Queries) GetOwner(ctx context.Context, id int64) (Owner, error) {
	row := q.db.QueryRow(ctx, getOwner, id)
	var i Owner
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
