package main

import (
    "net/http"
    "crypto/rand"
    "math/big"
    "time"
    "fmt"
    "html/template"
)

type tempContents struct {
    RandNum string
    Name    string
}

func main() {
    serve("9021")
}

func serve(port string) {
    http.HandleFunc("/", index)
    http.HandleFunc("/index.html", index)
    http.ListenAndServe(":" + port, nil)
}

func randInt(m int) (r string) {
    bigInt := big.NewInt(int64(m))
    num, err := rand.Int(rand.Reader, bigInt)
    if err != nil {
        fmt.Println(err)
    }
    numStr := num.String()
    return numStr
}

func index(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    randomInteger := randInt(5000)
    name := r.FormValue("name")
    if name == "" {
        name = "Nameless"
    }
    contents := tempContents{
        RandNum: randomInteger,
        Name: name,
    }
    tmpl, err := template.ParseFiles("html/index.html")
    if err != nil {
        panic(err)
    }
    fmt.Println("[200] OK " + name + ", " + randomInteger)
    duration := time.Since(start) * 100
    fmt.Println(duration)
    tmpl.Execute(w, contents)
}