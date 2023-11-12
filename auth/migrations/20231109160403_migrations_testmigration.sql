-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    name varchar(255) not null,
    email varchar(255) not null unique,
    role varchar(32),
    id serial not null unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd