-- +goose Up
CREATE TABLE concerts (
    id integer PRIMARY KEY CHECK (id > 0),
    title text NOT NULL CHECK (btrim(title) <> ''),
    description text NOT NULL CHECK (btrim(description) <> ''),
    capacity integer NOT NULL,
    available integer NOT NULL,
    CONSTRAINT concerts_available_check CHECK (available >= 0 AND available <= capacity)
);

INSERT INTO concerts (id, title, description, capacity, available)
VALUES (1, 'Вечер джаза', 'Концерт в пятницу', 10, 10);

-- +goose Down
DROP TABLE concerts;
