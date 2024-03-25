CREATE TABLE users
(
    id        SERIAL PRIMARY KEY,
    email     VARCHAR(255) UNIQUE NOT NULL,
    pass_hash BYTEA               NOT NULL
);
