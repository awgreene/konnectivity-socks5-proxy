package main

import (
	socks5 "github.com/armon/go-socks5"
)

type SimpleLogger interface {
	Println(v ...interface{})
}

type Foo struct {
	Logger SimpleLogger
}

func (f *Foo) Println(v ...interface{}) {
	f.Logger.Println(v...)
}

func main() {
	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "localhost:8080"); err != nil {
		panic(err)
	}
}
