package main

/*
NEXT STEPS
1. session
2. postgres
3. CRUDs
4. login/logout
5. authentication
6. go-demeler
*/

import (
	"fmt"
  "net"
	"net/http"
  "database/sql"
  _ "github.com/lib/pq"
)

const (
  host = "localhost"
  port = 5432
  user = "godev"
  password = "coco"
  dbname = "mcps"
)

type Application struct {
	mux *http.ServeMux
  db *sql.DB
}

func (a *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  host, port, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Printf("Hello, this is app router\n\n%+v\n\napp route host: '%s', port: %s\n\n", r, host, port)
	a.mux.ServeHTTP(w, r)
}

func (a *Application) crap(w http.ResponseWriter, r *http.Request) {
//------------------------------------------------------------------------------------------------------------
  sqlStatement := `SELECT username, password FROM users WHERE username=$1`
  var username string
  var password string
  row := a.db.QueryRow(sqlStatement, "mike")
  switch err := row.Scan(&username, &password); err {
  case sql.ErrNoRows:
    fmt.Println("No rows were returned!")
  case nil:
    fmt.Println(username, password)
  default:
    panic(err)
  }
//------------------------------------------------------------------------------------------------------------
	fmt.Fprintf(w, "Hello, this is crap\n")
}

func welcome(w http.ResponseWriter, r *http.Request) {
  switch r.RequestURI {
  case "/":
    fmt.Fprintf(w, "This is the welcome page\n\n")
  default:
    fmt.Fprintf(w, "This is an attempted hack\n\n")
  }
}

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  err = db.Ping()
  if err != nil {
    panic(err)
  } else {
    var res sql.Result
    res, err = db.Exec("SET search_path TO development")
    if err != nil {
      panic(err)
    }
    fmt.Printf("db res-->%v, err-->%v\n\n", res, err)
  }

	app := &Application{ mux: http.NewServeMux(), db: db }
	app.mux.HandleFunc("/crap", app.crap)
  app.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
  app.mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./public/images"))))
  app.mux.HandleFunc("/", welcome)
	http.ListenAndServe(":9292", app)
}
