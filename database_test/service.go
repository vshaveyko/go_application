package database_test

import (
  "fmt"
  "server/model"
  "server/db"
  "strconv"
)

func Test(number string) {

  db := db.Connection()

  var len int;

  if s, err := strconv.ParseInt(number, 10, 32); err == nil {
    len = int(s)
  }

  for i := 0; i < len; i++ {
    db.Create(&model.User{
      Name:        fmt.Sprintf("root %d", i),
      Description: fmt.Sprintf("description %d", i),
      PhoneNumber: fmt.Sprintf("+%d%d%d-%d%d%d-%d%d%d%d", i, i, i, i+1, i+1, i+1, i+2, i+2, i+2, i+2),
    })
  }

  var users []model.User

  db.Model(&users).Select()

  for _, value := range users {
    db.Delete(&value)
  }

}
