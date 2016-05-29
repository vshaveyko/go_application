package db

import (
    "gopkg.in/pg.v4"
)

var db = pg.Connect(&pg.Options{
            User: "postgres",
            Password: "root",
            Database: "go_test_dev"});

func Connection() *pg.DB {
  return db;
}
