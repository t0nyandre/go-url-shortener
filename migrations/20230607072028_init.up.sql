create table urls (
    id serial primary key,
    short_url text not null,
    long_url text not null,
    clicks int not null default 0,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);