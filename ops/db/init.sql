create table if not exists "user"
(
    id         serial,
    email      varchar(255),
    password   varchar(255),
    first_name varchar(100),
    last_name  varchar(100),
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (id)
);

create table if not exists book
(
    id         serial,
    title      varchar(255),
    author     varchar(255),
    pages      integer,
    quantity   integer,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (id)
);

create table if not exists book_user
(
    user_id    varchar(50),
    book_id    varchar(50),
    created_at timestamp,
    PRIMARY KEY (user_id, book_id)
);