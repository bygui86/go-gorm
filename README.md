
# go-gorm

Sample project to explore gorm Golang library for SQL

## TODO

- [x] simple example
    - [x] sqlite
    - [x] mysql
    - [x] postgres
- [x] relation example
    - [x] sqlite
    - [x] mysql
    - [x] postgres
- [x] transaction error handling
- [x] soft delete / delete
- [x] constraints
- [x] preload
- [ ] expose rest api
- [ ] prometheus metrics

## Run

### SQLite sample

```shell script
make run DB_TYPE=sqlite
```

### MySQL sample

```shell script
make run-mysql
make run DB_TYPE=mysql
```

### PostgreSQL sample

```shell script
make run-postgres
make run DB_TYPE=postgres
```

## Links

- https://gorm.io/index.html
- https://github.com/chzyer/logex
