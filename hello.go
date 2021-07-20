package main

import (
 "flag"
 "fmt"
 "log"
 "net/http"
)

var PORT = flag.Int("port", 15001, "listen port")

func Hello(w http.ResponseWriter, r *http.Request) {

 keys, ok := r.URL.Query()["name"]

 if !ok || len(keys[0]) < 1 {
  log.Println("Url параметр 'name' пропущен")
  return
 }

 // Query()["key"] gives us an array,
 //but we need only one element in it
 key := keys[0]

 log.Println("Url параметр 'name': " + string(key))

 message := "Hello, " + string(key) + "!"
 head := ""
 response := "" + head + "" + message + ""

 w.Header().Set("Content-Type", "text/html")
 w.Write([]byte(response))
}

func main() {
 flag.Parse()
 addr := fmt.Sprintf(":%d", *PORT)
 log.Println("Программа запущена на порту " + addr)
 http.HandleFunc("/hello", Hello)
 err := http.ListenAndServe(addr, nil)
 if err != nil {
  log.Fatal("ListenAndServe: ", err)
 }
}
