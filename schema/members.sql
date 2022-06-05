CREATE TABLE public.members
(
    id VARCHAR(20) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    create_time timestamp not null default now(),
    update_time timestamp
);
