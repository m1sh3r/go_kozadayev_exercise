package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Car struct {
	ID             int64
	Brand          string
	Model          string
	Year           int32
	EngineID       pgtype.Int8
	TransmissionID pgtype.Int8
}

type Engine struct {
	ID     int64
	Name   string
	Power  int32
	Volume float32
}

type Transmission struct {
	ID   int64
	Type string
}
