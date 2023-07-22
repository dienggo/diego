-- +goose Up
-- +goose StatementBegin
create table settings
(
    id         int auto_increment,
    `key`      varchar(255)                        null,
    value      text                                null,
    created_at timestamp default current_timestamp null,
    updated_at timestamp default current_timestamp null,
    constraint settings_pk
        primary key (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS settings
-- +goose StatementEnd
