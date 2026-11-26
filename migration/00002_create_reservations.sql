-- +goose Up
CREATE TABLE reservations (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    concert_id integer NOT NULL REFERENCES concerts (id),
    buyer_id text NOT NULL,
    key text NOT NULL,
    status text NOT NULL CHECK (status IN ('held', 'confirmed', 'expired')),
    expires_at timestamptz NOT NULL,
    CONSTRAINT reservations_buyer_key_unique UNIQUE (buyer_id, key)
);

-- +goose Down
DROP TABLE reservations;
