-- +goose Up
-- +goose StatementBegin
ALTER TABLE products
    DROP COLUMN IF EXISTS price,
    DROP COLUMN IF EXISTS quantity;
ALTER TABLE product_entries
    DROP COLUMN IF EXISTS product_full_name,
    DROP COLUMN IF EXISTS weight,
    DROP COLUMN IF EXISTS availability;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products
    ADD COLUMN price NUMERIC(10, 2) CHECK(price >= 0) NOT NULL,
    ADD COLUMN quantity INTEGER CHECK(quantity >= 0) NOT NULL;
ALTER TABLE product_entries
    ADD COLUMN product_full_name VARCHAR(255) UNIQUE NOT NULL,
    ADD COLUMN weight BIGINT,
    ADD COLUMN availability BOOLEAN;
-- +goose StatementEnd
