package main

import (
	"log"
	"net/http"
	"time"
)

func main()  {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello, world!"))
	})
	go http.ListenAndServe(":8080", nil)

	for {
		log.Println("service is running...")
		time.Sleep(5 * time.Second)
	}
}
