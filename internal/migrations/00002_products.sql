

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    -- user_id
    image BYTEA,
    display_name VARCHAR(50) UNIQUE  NOT NULL,
    quantity INTEGER CHECK(quantity >= 0) NOT NULL,
    price NUMERIC(10, 2) CHECK(price >= 0) NOT NULL,
    rating INTEGER CHECK(rating >= 1 and rating <= 5),
    description TEXT,
    category VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    activation boolean
)

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
