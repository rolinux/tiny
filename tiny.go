package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

// parse -port as string as we need a string for ListenAndServe and WriteString
// no validation of port string (maybe later)
var portPtr = flag.String("port", "8080", "port number to listen to")

func tiny(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost"
	}
	output := "<html><body><center>Welcome to <b>rolinux/tiny</b> app running<br>"
	output = output + "on <b>" + hostname + ":" + *portPtr + "</b> (hostname:port)<br>on <b>"
	output = output + runtime.GOOS + ":" + runtime.GOARCH + "</b> architecture</center></body></html>"
	io.WriteString(w, output)
}

func main() {
	runtime.GOMAXPROCS(1)
	flag.Parse()
	http.HandleFunc("/", tiny)
	// it will fail if the command-line flag is not a port number
	// not the most efficient string concat but speed is not the target
	if err := http.ListenAndServe(":"+*portPtr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
