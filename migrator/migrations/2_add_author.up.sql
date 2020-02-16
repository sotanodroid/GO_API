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
