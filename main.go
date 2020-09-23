package main

import (
    "net/http"
    "math/rand"
    "time"
    "fmt"
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

func main() {
    // create a new handler
    handler := HttpHandler{}

    // listen and serve
    http.ListenAndServe(":9000", handler)
}
