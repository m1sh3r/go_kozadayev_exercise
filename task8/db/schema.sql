CREATE TABLE cars (
    id BIGSERIAL PRIMARY KEY,
    brand TEXT NOT NULL
);

CREATE TABLE bookings (
    id BIGSERIAL PRIMARY KEY,
    car_id BIGINT NOT NULL REFERENCES cars(id),
    start_date BIGINT NOT NULL,
    end_date BIGINT NOT NULL
);
