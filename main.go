package main

import (
	"fmt"
	"github.com/shamaton/msgpack"
	"log"
)

func main() {
	v := "hello galesson."
	b, err := msgpack.Encode(v)
	if err != nil {
		log.Fatal(err)
	}

	var vv string
	err = msgpack.Decode(b, &vv)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(vv)
}

func a() {
	fmt.Println("a")
}

