package main

import (
	"os"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"
	"time"
)

var upTime = time.Now()

func handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	str := "<h1>It works!</h1>"

	env, exists := os.LookupEnv("ENV")
	if exists {
		str += fmt.Sprintf("env: %v<br>\n", env)
	}

	str += fmt.Sprintf("uptime: %v<br>\n", upTime.String())
	str += fmt.Sprintf("now: %v<br>\n", now.String())
	str += fmt.Sprintf("%v<br>\n", runtime.GOOS)
	str += fmt.Sprintf("%v<br>\n", runtime.GOARCH)
	str += fmt.Sprintf("%v<br>\n", runtime.Version())

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatalln(err)
	}
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Fatalln(err)
		}
		// handle err
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				str += fmt.Sprintf("%v<br>\n", v.IP)
			case *net.IPAddr:
				str += fmt.Sprintf("%v<br>\n", v.IP)
			}
		}
	}

	fmt.Fprintf(w, str)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
