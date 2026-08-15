package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("⚡ Orders Service running on :8081")
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"status\":\"ok\",\"orders\":[]}"))
	})
	_ = http.ListenAndServe(":8081", nil)
}
