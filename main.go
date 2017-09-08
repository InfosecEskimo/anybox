package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/franela/goreq"
)

const (
  port = ":8080"
)

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

func get_content(w http.ResponseWriter, r *http.Request) {
    // json data
    url := "http://localhost:9200/new-catalog/_search"

    res, err := http.Get(url)

    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(err.Error())
    }

    data := "" //tracks
    json.Unmarshal(res.Body, data)
    fmt.Printf("Results: %v\n", data)
    //os.Exit(0)
}

func init() {
  fmt.Printf("Started server at http://localhost%v.\n", port)
  http.HandleFunc("/", get_content)
  http.ListenAndServe(port, nil)
}

func main() {
  //get_content()
}
