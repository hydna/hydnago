package main

import "fmt"
import "bytes"
import "log"
import "github.com/hydna/hydnago"

func main() {
    data := bytes.NewBufferString("Hello world from GO!")

    err := hydna.Send("testing.hydna.net/", data)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully pushed data to url");
}