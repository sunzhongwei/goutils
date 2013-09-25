// > go run hello.go 
// or 
// > go build hello.go
// > ./hello

package goutils 

import (
    "testing"
)


func TestSayHello(t *testing.T) {
    const in = "Zhongwei"
    const out = "Hello world! Zhongwei"
    if x := SayHello(in); x != out {
        t.Errorf("SayHello(%v) = %v, want %v", in, x, out)
    }
}

