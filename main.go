package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "index.html") })

	fmt.Println("[INFO] Starting the server at 0.0.0.0:80")

	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		fmt.Println("[ERROR] Shutting down the server...")

	}
}
