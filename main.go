package main

import (
    "net/http"
    "fmt"
    "encoding/json"
    "strconv"

    qrcode "github.com/skip2/go-qrcode"
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
    // limit the uploaded file size to 10mbs
    request.ParseMultipartForm(10 << 20)
    var size, content string = request.FormValue("size"), request.FormValue("content")
    var codeData []byte

    writer.Header().Set("Content-Type", "application/json")

    // incase our content is not passed in or invalid through an error
    if content == "" {
        writer.WriteHeader(400)
        json.NewEncoder(writer).Encode("Could not determine the desired QRcode content!")
        return
    }

    // incase our size is not passed in or invalid through an error
    qrCodeSize, err := strconv.Atoi(size)
    if err != nil || size == "" {
        writer.WriteHeader(400)
        json.NewEncoder(writer).Encode("Could not determine the QRcode size!")
        return
    }

    qrCode := simpleQrcode{Content: content, Size: qrCodeSize}
    codeData, err = qrCode.generate()
    // incase we run into an error as we make the qrcode then thro an error back to the user as json
    if err != nil {
        writer.WriteHeader(400)
        json.NewEncoder(writer).Encode(fmt.Sprintf("Could not generate QR code. %v", err))
        return
    }

    writer.Header().Set("Content-Type", "image/png")
    writer.Write(codeData)
}

// main function to run the web server
func main(){
    http.HandleFunc("/generate", handleRequest)
    fmt.Println("[*] Server Running on port 8000...")
    http.ListenAndServe(":8000", nil)
}
