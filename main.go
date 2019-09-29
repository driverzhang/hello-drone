package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_,err := fmt.Fprintf(w, "您看到我了! CI / CD 成功了！！！")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	log.Println("http listen start port 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("List 9090")
	}
}
