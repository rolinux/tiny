package main

import (
        "io"
        "log"
        "flag"
        "net/http"
)

// parse -port as string as we need a string for ListenAndServe and WriteString
var portPtr = flag.String("port", "8000", "port number to listen to")

func port(w http.ResponseWriter, r *http.Request) {
          io.WriteString(w, *portPtr + "\n")
}

func main() {
        flag.Parse()
        http.HandleFunc("/", port)
        // it will fail if the command-line flag is not a port number
        // not the most efficient string concat but speed is not the target
        // http.ListenAndServe(":" + *portPtr, nil)
        if err := http.ListenAndServe(":" + *portPtr, nil);err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}
