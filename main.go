package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/iggyray/go-server/internal/service"
)

func main() {
	http.HandleFunc("/post", postHandler)

	log.Printf("Listening on http://localhost:%d\n", 8000)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	postService := service.NewPostService()

	post := postService.GetPost()

	log.Printf("%+v", post)

	_, err := io.WriteString(w, fmt.Sprintf("%+v", post))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
