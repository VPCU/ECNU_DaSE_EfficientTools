package main


import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "今天我要吃蛋饼\n(っˆڡˆς)\n\tVPCU")
}


func main() {
    http.HandleFunc("/", helloHandler)


    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
