package main

import(
  /*
  "log"
  "net/http"
  "path/filepath"
  "sync"
  "text/template"
  */
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "./handler"
)

/*
type templateHandler struct {
  once sync.Once
  filename string
  templ *template.Template
}


func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  t.once.Do(func() {
    t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
  })
  t.templ.Execute(w, nil)
}
*/

func main() {
  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/hello", handler.MainPage())
  e.Start(":8080")
  /*
  http.Handle("/", &templateHandler{filename: "chat.html"})
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal("ListenAndServe", err)
  }
  */
}
