create table urls (
    id serial primary key,
    short_url text not null,
    long_url text not null,
    clicks int not null default 0,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

-- create or replace function update_updated_at()
-- returns trigger as $$
-- begin
--     new.updated_at = now();
--     return new;
-- end;
-- $$ language plpgsql;

-- create trigger update_urls_updated_at
-- before update on urls
-- for each row
-- execute function update_updated_at();
