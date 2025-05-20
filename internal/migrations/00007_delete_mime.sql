-- +goose Up
-- +goose StatementBegin
ALTER TABLE products 
    DROP COLUMN IF EXISTS mime_type;

ALTER TABLE users
    DROP COLUMN IF EXISTS image;

ALTER TABLE users
    ADD COLUMN IF NOT EXISTS image_url VARCHAR(255);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
ALTER TABLE products
    ADD COLUMN IF NOT EXISTS mime_type VARCHAR(50);

ALTER TABLE users
    ADD COLUMN IF NOT EXISTS image BYTEA;

ALTER TABLE users
    DROP COLUMN IF EXISTS image_url;
-- +goose StatementEnd
