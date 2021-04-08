# Web API Demo with Golang

```bash
# build images
$ docker-compose build

# setup db schema
$ docker-compose run --rm db_migration goose up
$ docker-compose run --rm db_migration goose status

# start web server
$ docker-compose up go_app

# run test
$ docker-compose run --rm db_migration goose -env test up
$ docker-compose run --rm db_migration goose -env test status
$ docker-compose run --rm go_app go test -v -race -cover  ./... -parallel 4
```
