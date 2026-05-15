package repository

type Booking struct {
	ID        int64
	CarID     int64
	StartDate int64
	EndDate   int64
}

type Car struct {
	ID    int64
	Brand string
}
