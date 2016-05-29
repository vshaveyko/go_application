package model

import(
  "fmt"
)

type User struct {
  Id     int64
  Name   string
  PhoneNumber string
  Description string
}

func (u User) String() string {
    return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.PhoneNumber)
}
