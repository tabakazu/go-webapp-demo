# Web API Demo with Golang

### Setup Docker Container
```bash
$ docker-compose build
```

### Setup Database
for MySQL
```bash
$ docker-compose run --rm goapp mysql -uroot -hmysql
mysql >
# Use database
USE golang_webapi_demo_dev;
# Create `items` table
CREATE TABLE `items` (
  `id` BINARY(16) NOT NULL,
  `title` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'タイトル',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='アイテム';
```

for PostgreSQL
```bash
$ docker-compose run --rm goapp psql -hpostgres -Uroot -dgolang_webapi_demo_dev
# Create `items` table
CREATE TABLE items (
  id uuid NOT NULL,
  title VARCHAR NOT NULL DEFAULT '',
  PRIMARY KEY (id)
);
```

### Stating Web Server
```bash
$ docker-compose run --rm --service-ports goapp go run main.go
```
