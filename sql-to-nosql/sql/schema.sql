create table if not exists activities
(
    id varchar primary key,
    user_id varchar,
    timestamp timestamp,
    data jsonb
);