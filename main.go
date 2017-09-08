package main

import (
  "fmt"
  "net/http"

  "github.com/urfave/negroni"
  "github.com/hoisie/mustache"
)

const (
  port = ":8080"
  elkPort = ":9200"
  elkStack = "localhost"
  action = "_search"
)

var calls = 0
/*
type Token struct {
    Raw       string                 // The raw token.  Populated when you Parse a token
    Method    SigningMethod          // The signing method used or to be used
    Header    map[string]interface{} // The first segment of the token
    Claims    Claims                 // The second segment of the token
    Signature string                 // The third segment of the token.  Populated when you Parse a token
    Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
}

func handler(w http.ResponseWriter, r *http.Request) {
    return
}
*/

func elk_query(w http.ResponseWriter, r *http.Request) {
  calls++
  fmt.Fprintf(w, "Hello, world! You have called me %d times.\n", calls)
}

func init() {
  fmt.Printf("Started server at http://localhost%v.\n", port)
  http.HandleFunc("/", elk_query)
  http.ListenAndServe(port, nil)
}
/*
func auth() {
    tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return hmacSampleSecret, nil
    })
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        //add elk_query here
        elk_query
        fmt.Println(claims["foo"], claims["nbf"])
    } else {
        fmt.Println(err)
      }
}
*/

func main() {}

/*
{
  r := mux.NewRouter()
  // Routes consist of a path and a handler function.
  r.HandleFunc("/", YourHandler)

  // Bind to a port and pass our router in
  log.Fatal(http.ListenAndServe(port, r))
}
*/
