// > go run hello.go 
// or 
// > go build hello.go
// > ./hello

package goutils 

import "fmt"

func SayHello(name string) string {
    msg := "Hello world! " + name
    fmt.Println(msg)
    return msg
}

