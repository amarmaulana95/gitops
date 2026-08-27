package main

import (
    "fmt"
    "net/http"
)

func healthz(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "OK")
}

func main() {
    http.HandleFunc("/healthz", healthz)

    fmt.Println("Backend Go CI/CD Running on :4000!")

    if err := http.ListenAndServe(":4000", nil); err != nil {
        panic(err)
    }
}
