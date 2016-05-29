package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/cors"
  "server/routes"
)

func main() {

  m := martini.Classic()

  m.Use(cors.Allow(&cors.Options{
    AllowOrigins:     []string{"*"},
    AllowMethods:     []string{"GET"},
    AllowHeaders:     []string{"Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
  }))

  routes.Include(m)
  m.Run();

}
