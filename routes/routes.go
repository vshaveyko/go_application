package routes

import (
  "github.com/go-martini/martini"

  "server/views"
  "server/database_test"
  "server/calculate_sum"

  "net/http"
)

func Include(m *martini.ClassicMartini) {

  m.Get("/", views.Home);

  m.Get("/calculate_sum", func(req *http.Request) string {
    if len(req.URL.Query()["number"]) > 0 {
      return calculate_sum.Calculate(req.URL.Query()["number"][0]);
    }

    return "Invalid params"
  });

  m.Get("/calculate_sum_routine", func(req *http.Request) string {
    if len(req.URL.Query()["number"]) > 0 {
      return calculate_sum.CalculateRoutine(req.URL.Query()["number"][0]);
    }

    return "Invalid params"
  });

  m.Get("/database_test", func(req *http.Request) {
    if len(req.URL.Query()["number"]) > 0 {
      database_test.Test(req.URL.Query()["number"][0]);
    }
  });

}
