package service

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

func (s *Service) All(res http.ResponseWriter, req *http.Request) {

}

func (s *Service) Init() {
	s.Get("/", http.HandlerFunc(s.All))
}
