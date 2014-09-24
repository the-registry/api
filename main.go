package main

import "versionsio/api/pkg/service"
import "github.com/lib/pq"
import "github.com/segmentio/go-log"
import "github.com/segmentio/go-env"
import "database/sql"

var Version = "0.0.1"

func main() {
	db, err := sql.Open("postgres", env.MustGet("VERSIONSIO_POSTGRES_URI"))
	log.Check(err)

	o := &service.Options{
		Db: db,
	}

	s := Service.New(o)
	s.Init()
	err = s.Listen()
	log.Check(err)
}
