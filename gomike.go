package main

import (
	"fmt"
	"net/http"
	"web"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat! mike 2")
}

func main() {
	//http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)

	http.HandleFunc("/secret", web.Secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

}
