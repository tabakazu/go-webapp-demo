# Web API Demo with Golang

```bash
# build images
$ docker-compose build

# setup db schema
$ docker-compose run --rm db_migration goose up
$ docker-compose run --rm db_migration goose status

# start web server
$ docker-compose up go_app
```
