package server

import (
	"fmt"
	"net/http"
)

func Launch(serverName string) {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm %s.", serverName)
	})

	http.HandleFunc("/" + serverName + "/hello", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "[%s] hello", serverName)
	})

	http.HandleFunc("/" + serverName + "/hey", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "[%s] hey", serverName)
	})

	fmt.Printf("launch %s", serverName)
  	http.ListenAndServe(":8080", nil)
}