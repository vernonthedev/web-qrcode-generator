package main

import (
    "net/http"
    "fmt"
)

//handle what ever request is sent under the generate route
func handleRequest(writer http.ResponseWriter, request *http.Request) {
    
}

// main function to run the web server
func main(){
    http.HandleFunc("/generate", handleRequest)
    fmt.Println("[*] Server Running on port 8000...")
    http.ListenAndServe(":8000", nil)
}
