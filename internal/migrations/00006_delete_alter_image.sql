-- +goose Up
-- +goose StatementBegin
ALTER TABLE products
    DROP COLUMN IF EXISTS image,
    ADD COLUMN IF NOT EXISTS image_url VARCHAR(255)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products
    ADD COLUMN IF NOT EXISTS image BYTEA,
    DROP COLUMN IF EXISTS image_url;
-- +goose StatementEnd
