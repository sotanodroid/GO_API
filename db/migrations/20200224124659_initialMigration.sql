
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE SCHEMA IF NOT EXISTS goapi;

CREATE TABLE IF NOT EXISTS goapi.authors
(
    id serial NOT NULL PRIMARY KEY,
    firstname VARCHAR (50),
    lastname VARCHAR (50)
);

INSERT INTO goapi.authors
(
    firstname,
    lastname
)
VALUES
(
    'John',
    'Doe'
),
(
    'Steve',
    'Smith'
);

CREATE TABLE IF NOT EXISTS goapi.books
(
    id serial NOT NULL PRIMARY KEY,
    isbn VARCHAR (50),
    title VARCHAR (50),
    author INT REFERENCES goapi.authors ON DELETE CASCADE
);

INSERT INTO goapi.books
(
    isbn,
    title,
    author
)
VALUES
(
    '153223',
    'Book_One',
    (SELECT id FROM goapi.authors WHERE firstname='John' AND lastname='Doe' LIMIT 1)
),
(
    '153235',
    'Book Two',
    (SELECT id FROM goapi.authors WHERE firstname='Steve' AND lastname='Smith' LIMIT 1)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS goapi.books;
DROP TABLE IF EXISTS goapi.authors CASCADE;
DROP SCHEMA IF EXISTS goapi CASCADE;
