# Web API Demo with Golang

```bash
# build images
$ docker-compose build

# setup db schema
$ docker-compose run --rm db_migration goose -path .dbschema create AddUsersTable sql
$ docker-compose run --rm db_migration goose -path .dbschema up
$ docker-compose run --rm db_migration goose -path .dbschema status

# start web server
$ docker-compose up go_app

# run test
$ docker-compose run --rm db_migration goose -env test -path .dbschema up
$ docker-compose run --rm db_migration goose -env test -path .dbschema status
$ docker-compose run --rm go_app go test -v -race -cover  ./... -parallel 4
```
