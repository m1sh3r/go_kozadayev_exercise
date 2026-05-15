CREATE TABLE owners (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE cars (
    id BIGSERIAL PRIMARY KEY,
    brand TEXT NOT NULL,
    owner_id BIGINT NOT NULL REFERENCES owners(id)
);

CREATE TABLE service_records (
    id BIGSERIAL PRIMARY KEY,
    car_id BIGINT NOT NULL REFERENCES cars(id),
    description TEXT NOT NULL,
    date BIGINT NOT NULL
);
