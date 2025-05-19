

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product_entries (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    product_full_name VARCHAR(255) UNIQUE  NOT NULL,
    quantity INTEGER CHECK(quantity >= 0) NOT NULL,
    price NUMERIC(10, 2) CHECK(price >= 0) NOT NULL,
    review TEXT,
    weight BIGINT,
    warranty_period INTERVAL,
    availability boolean,
    rating INTEGER CHECK(rating >= 1 and rating <= 5),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT check_warranty_reasonable 
        CHECK(warranty_period IS NULL OR 
        (warranty_period >= INTERVAL '0 days' AND warranty_period <= INTERVAL '10 years')
        ),
    CONSTRAINT check_availability_quantity 
        CHECK(availability = FALSE OR quantity > 0),
    CONSTRAINT check_rating_has_review 
        CHECK(rating IS NULL OR review IS NOT NULL)
)

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE product_entries;
-- +goose StatementEnd