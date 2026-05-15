-- name: AssembleCar :one
UPDATE cars 
SET engine_id = $2, transmission_id = $3 
WHERE id = $1 
RETURNING *;

-- name: GetCarSpec :one
SELECT 
    c.id, c.brand, c.model, c.year,
    e.id as engine_id, e.name as engine_name, e.power as engine_power, e.volume as engine_volume,
    t.id as transmission_id, t.type as transmission_type
FROM cars c
LEFT JOIN engines e ON c.engine_id = e.id
LEFT JOIN transmissions t ON c.transmission_id = t.id
WHERE c.id = $1;
