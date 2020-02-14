CREATE TABLE IF NOT EXISTS goapi.Book
(
    id serial NOT NULL PRIMARY KEY,
    Isbn VARCHAR (50),
    Title VARCHAR (50),
    Author INT REFERENCES goapi.Author(id) ON DELETE CASCADE
);
