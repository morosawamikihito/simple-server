package server

import (
	"fmt"
	"net/http"
	"net"
)

func getIPAddress() []string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err.Error())
	}

	ipv4 := []string{}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipv4 = append(ipv4, ipnet.IP.String())
			}
		}
	}
	return ipv4
}

func Launch(serverName string) {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm %s.", serverName)
	})

	http.HandleFunc("/" + serverName + "/ip", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "[%s] ip: %s", serverName, getIPAddress())
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