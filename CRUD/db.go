package main

import (
	"database/sql"
	"log"
)

/*
docker run --name test-mysql -e MYSQL_ROOT_PASSWORD=strong_password -d mysql
docker exec -it mysql bash

# mysql -u root -p

mysql> create database devbook;
mysql> use devbook;
mysql> create table user(
    -> id int auto_increment primary key,
    -> name varchar(50) not null,
    -> email varchar(50) not null
    -> ) ENGINE=INNODB;
mysql> create user 'user'@'%' identified by 'pass';
mysql> GRANT ALL PRIVILEGES ON devbook.* TO 'user'@'%' WITH GRANT OPTION;
*/

const url = "user:pass@/devbook?charset=utf8&parseTime=True&loc=Local"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", url)

	if err != nil {
		log.Fatal("unable to connect to MySQL:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("unable to ping MySQL:", err)
		return nil, err
	}

	return db, nil
}
