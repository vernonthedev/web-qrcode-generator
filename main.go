package main

import (
    "net/http"
    "fmt"
)

//define the qrcode properties using a struct
type simpleQrcode struct{
    Content string
    Size    int
}

func (code *simpleQrcode) generate() ([]byte,error){
    // generate qrcode using go's inbuilt qrcode lib using size and the content frm the simpleQrcode struct
    qrCode, err := qrcode.Encode(code.Content, qrcode.Medium, code.Size)
    //incase we have any errors as we generate the qrcode then show the errors.
    if err != nil {
        return nil, fmt.Errorf("Could not generate qrcode: %v", err)
    }
    return qrCode, nil
}


//handle what ever request is sent under the generate route
func handleRequest(writer http.ResponseWriter, request *http.Request) {
    
}

// main function to run the web server
func main(){
    http.HandleFunc("/generate", handleRequest)
    fmt.Println("[*] Server Running on port 8000...")
    http.ListenAndServe(":8000", nil)
}
