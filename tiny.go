package main

import (
        "io"
        "os"
        "net/http"
)

func port(w http.ResponseWriter, r *http.Request) {
        hostname, err := os.Hostname()
        if err == nil {
          io.WriteString(w, hostname)
        }
}

func main() {
        http.HandleFunc("/", port)
        http.ListenAndServe(":8000", nil)
}
