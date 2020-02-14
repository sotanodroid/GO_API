# Simple GO Lang REST API

> Simple RESTful API to create, read, update and delete books. No database implementation yet

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
docker run -v { path_to_migration_folder }:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
```

``` bash
go build github.com/sotanodroid/GO_API/cmd/goapi
./goapi
```

Server would be awailable on port 8000

## Endpoints

### Get All Books
``` bash
GET api/books
```
### Get Single Book
``` bash
GET api/books/{id}
```

### Delete Book
``` bash
DELETE api/books/{id}
```

### Create Book
``` bash
POST api/books

# Request sample
# {
#   "isbn":"4545454",
#   "title":"Book Three",
#   "author":{"firstname":"Harry",  "lastname":"White"}
# }
```

### Update Book
``` bash
PUT api/books/{id}

# Request sample
# {
#   "isbn":"4545454",
#   "title":"Updated Title",
#   "author":{"firstname":"Harry",  "lastname":"White"}
# }

```
