package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/franela/goreq"
  "github.com/joho/godotenv"
  "log"
  "os"
)

type envObj struct {
  PORT string
  URL string
}

func loadEnv() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file, see .env.example")
    return false
  }
  env := &envObj{
    PORT = os.Getenv("PORT"),
    URL = os.Getenv("URL")
  }

  return env
}

func elk_search(w http.ResponseWriter, r *http.Request) {
  res, err := goreq.Request{
  	Method:      "POST",
  	Uri:         "http://localhost:9200/new-catalog/_search",
  	Compression: goreq.Gzip(),
  	ShowDebug:   true,
  }.Do()
  // res.Response will contain the original http.Response structure
  fmt.Println("<h1>Title: Test Site</h1>")
  fmt.Println(res.Response, err)
  //return res.Response
}

func init() {
  fmt.Printf("Started server at http://localhost%v.\n", port)
  http.HandleFunc("/", get_content)
  http.ListenAndServe(port, nil)
}

func main() {
  env = loadEnv()
}
