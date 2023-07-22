-- +goose Up
-- +goose StatementBegin
create table users
(
    id         int auto_increment,
    name       varchar(255)                        null,
    email      varchar(255)                        null,
    created_at timestamp default current_timestamp null,
    updated_at timestamp default current_timestamp null,
    constraint settings_pk
        primary key (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users
-- +goose StatementEnd
