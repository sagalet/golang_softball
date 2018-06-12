package main

import (
	"content1"
	"html/template"
	"log"
	"net/http"
	"packages/config"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
)

var RES_PATH = "/home/davis_chen/golang_softball/res"
var HTTP_PATH = "/home/davis_chen/golang_softball/html"
var JS_PATH = "/home/davis_chen/golang_softball/js"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("index.html")
		t, _ = t.ParseFiles(HTTP_PATH + "/index.html")
		t.Execute(w, nil)
	})
	http.HandleFunc("/pic/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", RES_PATH+r.URL.Path[0:])
		http.ServeFile(w, r, RES_PATH+r.URL.Path[0:])
	})

	http.HandleFunc("/js/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", JS_PATH+r.URL.Path[3:])
		http.ServeFile(w, r, JS_PATH+r.URL.Path[3:])
	})

	http.HandleFunc("/json/", func(w http.ResponseWriter, r *http.Request) {
		c := config.GetConfig("config/default.json")
		b := config.GetByteArray(c)
		log.Println("send json object")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)
	})

	http.HandleFunc("/proto/", func(w http.ResponseWriter, r *http.Request) {
		d := content1.Pack1{
			Data: []*content1.Data1{
				{
					Name: "user1",
					Id:   1,
					Msg:  "just user1"},
				{
					Name: "user2",
					Id:   2,
					Msg:  "this is user2"},
			},
		}

		out, err := proto.Marshal(&d)
		if err != nil {
			log.Fatalln("failed to marshal: ", err)
		}
		log.Printf("data=%s", out[0:])
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(out)
	})

	http.HandleFunc("/redis/", func(w http.ResponseWriter, r *http.Request) {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		val, err := client.Get("name").Result()
		if err != nil {
			panic(err)
		}
		log.Println("name=", val)
	})

	log.Println("start server!!")

	log.Fatal(http.ListenAndServe(":80", nil))
}
