package main

import (
    "net/http"
    "crypto/rand"
    "math/big"
    "time"
    "fmt"
    "html/template"
)

func main() {
    serveSlash()
}

func serveSlash() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/index.html", hello)
    http.ListenAndServe(":9021", nil)
}

type tempContents struct {
    RandNum string
    Name    string
}

func hello(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    bigInt := big.NewInt(int64(5000))
    num, err := rand.Int(rand.Reader, bigInt)
    if err != nil {
        fmt.Println(err)
    }
    numStr := num.String()
    respString := r.FormValue("name")
    if respString == "" {
        respString = "Nameless"
    }
    contents := tempContents{
        RandNum: numStr, 
        Name: respString,
    }
    tmpl, err := template.ParseFiles("html/index.html")
    if err != nil {
        panic(err)
    }
    fmt.Println("[200] OK " + respString + ", " + numStr)
    duration := time.Since(start) * 100
    fmt.Println(duration)
    tmpl.Execute(w, contents)
}