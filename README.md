# Web API Demo with Golang

```bash
# build images
$ docker-compose build

# setup db schema
$ docker-compose run --rm db_migration goose create CreateItemsTable sql
$ docker-compose run --rm db_migration goose up
$ docker-compose run --rm db_migration goose status

# login mysql
$ docker-compose run --rm db_migration mysql -uroot -hmysql

# start web server
$ docker-compose up go_app

# run test
$ docker-compose run --rm db_migration goose -env test up
$ docker-compose run --rm db_migration goose -env test status
$ docker-compose run --rm go_app go test -v -race -cover  ./... -parallel 4

# run production image from local
$ docker build -t go-webapp .
$ docker run -p 8080:8080 -e MYSQL_URL='root:@tcp(host.docker.internal:3306)/golang_webapi_demo_dev' --name go-webapp go-webapp
$ docker rm go-webapp

# generate api doc
$ docker-compose run --rm go_app swag init -o .docs/api -g interfaces/webapi/server.go
```

# Directory Structure
### Domain (Enterprise Business Rules)
- domain
  - entity
  - value
  - service
  - enum
  - collection

### Application (Application Business Rules)
- application
  - service

### Interface Adapters
- controller
  - interfaces/webapi/controller (webapi)
- gateway
  - external/datastore/repository (datastore)
- presenter

### Frameworks & Drivers
- user interfaces
  - interfaces/webapi (webapi)
- external interfaces
  - external/datastore (datastore)
