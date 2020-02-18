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
