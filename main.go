package main

import "github.com/the-registry/api/pkg/service"
import "github.com/segmentio/go-log"
import elastigo "github.com/mattbaird/elastigo/lib"
import "fmt"

var Version = "0.0.1"

func main() {
	o := &service.Options{
		Db: elastigo.NewConn(),
	}

	s := service.New(o)
	s.Init()

	u := "localhost:3000"

	fmt.Printf("Running server at %s\n", u)

	err := s.Listen(u)
	log.Check(err)
}
