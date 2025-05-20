-- +goose Up
-- +goose StatementBegin
ALTER TABLE products DROP CONSTRAINT IF EXISTS products_display_name_key;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products ADD CONSTRAINT products_display_name_key UNIQUE (display_name);
-- +goose StatementEnd
