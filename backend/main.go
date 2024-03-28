package main

import (
	"fmt"
	// "io"
	"net/http"
	// "os"
)

func main() {
    mux := http.NewServeMux();
    mux.HandleFunc("/", serveFile);
    mux.HandleFunc("/api/hello", hello);
    err := http.ListenAndServe(":6969", mux);

    if err != nil {
        fmt.Printf("There was an error while running the server:\n%s", err.Error());
        return;
    }
}

const FRONTEND_ROOT  = "../frontend";
const INDEX_HTML_PATH = "/src/index.html";

func serveFile(w http.ResponseWriter, req *http.Request) {
    path := FRONTEND_ROOT + req.URL.Path;

    if req.URL.Path == "/" {
        path += INDEX_HTML_PATH;
    }

    fmt.Println("Received a request for path " + path);
    http.ServeFile(w, req, path);
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Hello");
}
