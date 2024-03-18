-- +goose Up
drop table if exists mobile_number;

create table mobile_number (
    mobile_number varchar(10) not null unique
);
-- +goose StatementBegin
SELECT 'up SQL query';

create or replace function prevent_duplicate_value()
returns trigger as $$
begin
    if exists (select 1 from mobile_number where mobile_number = new.mobile_number) then
        return null;
    end if;
    return new;
end;
$$ language plpgsql;

create trigger check_duplicate_value
before insert on mobile_number
for each row
execute function prevent_duplicate_value();

-- +goose StatementEnd

-- +goose Down
drop table if exists mobile_number;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
