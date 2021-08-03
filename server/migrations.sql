begin;

create table if not exists songs (
  id serial primary key,
  name text unique not null,
  labels jsonb
);

commit;
