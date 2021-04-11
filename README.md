# Web API Demo with Golang

```bash
# build images
$ docker-compose build

# setup db schema
$ docker-compose run --rm db_migration goose create AddUsersTable sql
$ docker-compose run --rm db_migration goose up
$ docker-compose run --rm db_migration goose status

# start web server
$ docker-compose up go_app

# run test
$ docker-compose run --rm db_migration goose -env test up
$ docker-compose run --rm db_migration goose -env test status
$ docker-compose run --rm go_app go test -v -race -cover  ./... -parallel 4

# run production image from local
$ docker build -t golang-webapi-demo-app .
$ docker run -p 8080:8080 -e MYSQL_URL='root:@tcp(host.docker.internal:3306)/golang_webapi_demo_dev' --name golang-webapi-demo-app golang-webapi-demo-app
$ docker rm golang-webapi-demo-app
```
