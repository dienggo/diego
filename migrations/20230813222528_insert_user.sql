-- +goose Up
-- +goose StatementBegin
insert into users (`name`, `email`, `password`, `token`)
values ("Daewu", "daewu.bintara1996@gmail.com", md5("password"), md5("token1234567890"));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
