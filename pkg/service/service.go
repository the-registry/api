package service

import "versionsio/api/routes"
import "github.com/gohttp/app"
import "database/sql/driver"
import "net/http"

type Options struct {
	Db *driver.Conn
}

type Service struct {
	*Options
	*app.App
}

func New(o *Options) *Service {
	return &Service{
		Options: o,
		App:     app.New(),
	}
}

func (s *Service) Init() {
	s.Get("/", http.HandlerFunc(routes.HomeHandler))
	s.Get("/:namespace", http.HandlerFunc(routes.IndexHandler))
	s.Get("/:namespace/search", http.HandlerFunc(routes.SearchHandler))
	// register
	s.Post("/:namespace", http.HandlerFunc(routes.IndexHandler))
}
