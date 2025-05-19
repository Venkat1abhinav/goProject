-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN mime_type  VARCHAR(50);
ALTER TABLE products ADD COLUMN mime_type  VARCHAR(50);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN mime_type;
ALTER TABLE products DROP COLUMN mime_type;
-- +goose StatementEnd
