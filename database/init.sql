DROP DATABASE IF EXISTS executiontimes_db;

CREATE DATABASE executiontimes_db;

\c executiontimes_db;

CREATE TABLE IF NOT EXISTS execution_times (
    id SERIAL PRIMARY KEY,
    test VARCHAR(255) NOT NULL,
    language VARCHAR(255) NOT NULL,
    algorithm VARCHAR(255) NOT NULL,
    time INTEGER NOT NULL
);