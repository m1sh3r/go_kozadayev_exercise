-- name: CreateOwner :one
INSERT INTO owners (name) VALUES ($1) RETURNING *;

-- name: CreateCar :one
INSERT INTO cars (brand, owner_id) VALUES ($1, $2) RETURNING *;

-- name: CreateServiceRecord :one
INSERT INTO service_records (car_id, description, date) VALUES ($1, $2, $3) RETURNING *;

-- name: GetOwner :one
SELECT * FROM owners WHERE id = $1;

-- name: GetDashboardData :many
SELECT 
    c.id as car_id, c.brand as car_brand,
    sr.id as service_id, sr.description as service_desc, sr.date as service_date
FROM cars c
LEFT JOIN (
    SELECT DISTINCT ON (car_id) *
    FROM service_records
    ORDER BY car_id, date DESC
) sr ON c.id = sr.car_id
WHERE c.owner_id = $1;
