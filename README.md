# Simple GOLang REST API

> Simple RESTful API to create, read, update and delete books.

## Quick Start


``` bash
# Make .env file
cp .env.example .env
```

```bash
# Run Docker-compose with PSQL
docker-compose up
```

```bash
# Run migrations
make migrate
```

``` bash
# Build app
go build github.com/sotanodroid/GO_API/cmd/goapi

#Or just run pre-built app
./main
```

Server would be awailable on port 8000

## Endpoints

### Get All Books
``` bash
GET api/books

# Response

[
    {
        "id": 1,
        "isbn": "153223",
        "title": "Book_One",
        "author": {
            "id": 1,
            "firstname": "John",
            "lastname": "Doe"
        }
    },
    {
        "id": 2,
        "isbn": "153235",
        "title": "Book Two",
        "author": {
            "id": 2,
            "firstname": "Steve",
            "lastname": "Smith"
        }
    }
]
```
### Get Single Book
``` bash
GET api/books/{id}

# Response
{
    "id": 1,
    "isbn": "153223",
    "title": "Book_One",
    "author": {
        "id": 1,
        "firstname": "John",
        "lastname": "Doe"
    }
}
```

### Delete Book
``` bash
DELETE api/books/{id}

Response:
"Deleted"
```

### Create Book
``` bash
POST api/books

# Request sample
{
  "isbn":"4545454",
  "title":"Book Three",
  "author": {
      "firstname":"John",
      "lastname":"Doe"
    }
}
```

### Update Book
``` bash
PUT api/books/{id}

# Request sample
{
  "isbn":"4545454",
  "title":"Updated Title"
}

```

#### Tests:

Start DataBase in Docker container, run migrations.
``` bash
go test ./... -v
```
