package gomike

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat! mike 2")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
