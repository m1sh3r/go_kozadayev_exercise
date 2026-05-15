CREATE TABLE engines (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    power INTEGER NOT NULL,
    volume REAL NOT NULL
);

CREATE TABLE transmissions (
    id BIGSERIAL PRIMARY KEY,
    type TEXT NOT NULL
);

CREATE TABLE cars (
    id BIGSERIAL PRIMARY KEY,
    brand TEXT NOT NULL,
    model TEXT NOT NULL,
    year INTEGER NOT NULL,
    engine_id BIGINT REFERENCES engines(id),
    transmission_id BIGINT REFERENCES transmissions(id)
);
