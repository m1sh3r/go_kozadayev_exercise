-- name: GetCarForUpdate :one
SELECT * FROM cars WHERE id = $1 FOR UPDATE;

-- name: CheckOverlap :one
SELECT EXISTS (
    SELECT 1 FROM bookings 
    WHERE car_id = $1 AND start_date < $3 AND $2 < end_date
) as overlapped;

-- name: CreateBooking :one
INSERT INTO bookings (car_id, start_date, end_date) VALUES ($1, $2, $3) RETURNING *;
