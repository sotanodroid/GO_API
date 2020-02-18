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
# Build migrations container
docker build -t migrator -f migrator/Dockerfile .

# Run migrations
make migrate
```

``` bash
# Build app
go build github.com/sotanodroid/GO_API/cmd/goapi
# Run app
./goapi
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


TODO
<!-- 
### Delete Book
``` bash
DELETE api/books/{id}
``` -->

### Create Book
``` bash
POST api/books

# Request sample
{
  "isbn":"4545454",
  "title":"Book Three",
  "author": {
      "firstname":"Harry",
      "lastname":"White"
    }
}
```


TODO
<!-- ### Update Book
``` bash
PUT api/books/{id}

# Request sample
# {
#   "isbn":"4545454",
#   "title":"Updated Title",
#   "author":{"firstname":"Harry",  "lastname":"White"}
# }

``` -->
