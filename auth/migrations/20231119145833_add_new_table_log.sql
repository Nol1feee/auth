-- +goose Up
-- +goose StatementBegin
CREATE TABLE logs (
    user_id int,
    date_created date,
    date_deleted date
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table logs;
-- +goose StatementEnd
