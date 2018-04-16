package main

import (
    "net/http"
    "html/template"
    "log"
)

var RES_PATH = "/home/davis_chen/golang_softball/res"
var HTTP_PATH = "/home/davis_chen/golang_softball/html"
var JS_PATH = "/home/davis_chen/golang_softball/js"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        t := template.New("index.html");
        t,_ = t.ParseFiles(HTTP_PATH + "/index.html");
        t.Execute(w, nil);
    })
    http.HandleFunc("/pic/", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s", RES_PATH + r.URL.Path[0:])
        http.ServeFile(w, r, RES_PATH + r.URL.Path[0:])
    })

    http.HandleFunc("/js/", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s", JS_PATH + r.URL.Path[3:])
        http.ServeFile(w, r, JS_PATH + r.URL.Path[3:])
    })

    log.Println("start server");

    log.Fatal(http.ListenAndServe(":80", nil))
}
