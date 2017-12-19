# kafka-mariadb-demo

## Requirements
 - Go
 - MariaDB
 - Kafka

## Start MariaDB container

```sh
docker run --name mariadb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=passwd -d mariadb:10
```

## Start kafka container

//TODO: curretly just running locally on host machine

## Create topic
kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test

## Create tables
```sql
CREATE DATABASE demo;
USE demo;
CREATE TABLE messages
(
    id INT(11) NOT NULL AUTO_INCREMENT,
    message VARCHAR(100),
    CONSTRAINT id PRIMARY KEY (id)
);
```

## Run

Go to repo root and run:

```sh
go run consumer/consumer.go
```

```sh
go run producer/producer.go
```
