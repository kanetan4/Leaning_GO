package main

import (
    "context"
    "fmt"
    "ex3_hellokitex/kitex_gen/hello/helloservice"
    "github.com/cloudwego/kitex/client"
)

func main() {
    c, err := helloservice.NewClient("hello", client.WithHostPorts("127.0.0.1:8888"))
    if err != nil {
        panic(err)
    }

    resp, err := c.SayHello(context.Background(), "Test")
    if err != nil {
        panic(err)
    }

    fmt.Println("Response from server:", resp)
}
