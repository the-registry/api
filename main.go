package main

import "versionsio/api/pkg/service"
import "github.com/segmentio/go-log"
import elastigo "github.com/mattbaird/elastigo/lib"

var Version = "0.0.1"

func main() {
	o := &service.Options{
		Db: elastigo.NewConn(),
	}

	s := service.New(o)
	s.Init()
	err := s.Listen("localhost:3000")
	log.Check(err)
}
