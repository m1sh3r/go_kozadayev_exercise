package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const assembleCar = `-- name: AssembleCar :one
UPDATE cars 
SET engine_id = $2, transmission_id = $3 
WHERE id = $1 
RETURNING id, brand, model, year, engine_id, transmission_id
`

type AssembleCarParams struct {
	ID             int64
	EngineID       pgtype.Int8
	TransmissionID pgtype.Int8
}

func (q *Queries) AssembleCar(ctx context.Context, arg AssembleCarParams) (Car, error) {
	row := q.db.QueryRow(ctx, assembleCar, arg.ID, arg.EngineID, arg.TransmissionID)
	var i Car
	err := row.Scan(
		&i.ID,
		&i.Brand,
		&i.Model,
		&i.Year,
		&i.EngineID,
		&i.TransmissionID,
	)
	return i, err
}

const getCarSpec = `-- name: GetCarSpec :one
SELECT 
    c.id, c.brand, c.model, c.year,
    e.id as engine_id, e.name as engine_name, e.power as engine_power, e.volume as engine_volume,
    t.id as transmission_id, t.type as transmission_type
FROM cars c
LEFT JOIN engines e ON c.engine_id = e.id
LEFT JOIN transmissions t ON c.transmission_id = t.id
WHERE c.id = $1
`

type GetCarSpecRow struct {
	ID               int64
	Brand            string
	Model            string
	Year             int32
	EngineID         pgtype.Int8
	EngineName       pgtype.Text
	EnginePower      pgtype.Int4
	EngineVolume     pgtype.Float4
	TransmissionID   pgtype.Int8
	TransmissionType pgtype.Text
}

func (q *Queries) GetCarSpec(ctx context.Context, id int64) (GetCarSpecRow, error) {
	row := q.db.QueryRow(ctx, getCarSpec, id)
	var i GetCarSpecRow
	err := row.Scan(
		&i.ID,
		&i.Brand,
		&i.Model,
		&i.Year,
		&i.EngineID,
		&i.EngineName,
		&i.EnginePower,
		&i.EngineVolume,
		&i.TransmissionID,
		&i.TransmissionType,
	)
	return i, err
}
