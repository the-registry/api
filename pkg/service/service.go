package service

import "github.com/gohttp/app"
import "github.com/gohttp/response"
import "github.com/gohttp/logger"
import elastigo "github.com/mattbaird/elastigo/lib"
import "net/http"
import "log"
import "encoding/json"
import u "versionsio/api/pkg/utils"

type Package struct {
	Url  string `json:"url"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Options struct {
	Db *elastigo.Conn
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

func (s *Service) HomeHandler(res http.ResponseWriter, req *http.Request) {
}

func (s *Service) IndexHandler(res http.ResponseWriter, req *http.Request) {
	// t := req.URL.Query().Get(":type")
	pkgs, err := s.Db.Search("versions", "Package", nil, "*")
	if err != nil {
		log.Fatal(err)
	}
	response.JSON(res, pkgs)
}

func (s *Service) SearchHandler(res http.ResponseWriter, req *http.Request) {

}

func (s *Service) CreateHandler(res http.ResponseWriter, req *http.Request) {
	t := req.URL.Query().Get(":type")
	var j map[string]string
	d := json.NewDecoder(req.Body)
	err := d.Decode(&j)

	if err != nil {
		u.Error(res, err)
		return
	}

	s.Db.Index("versions", "Packages", "", nil, Package{
		Name: j["name"],
		Url:  j["url"],
		Type: t,
	})

	res.WriteHeader(http.StatusCreated)
}

func (s *Service) Init() {
	s.Use(logger.New())
	s.Get("/", http.HandlerFunc(s.HomeHandler))
	s.Get("/:type/packages", http.HandlerFunc(s.IndexHandler))
	s.Get("/:type/packages/search", http.HandlerFunc(s.SearchHandler))
	s.Post("/:type/packages", http.HandlerFunc(s.CreateHandler))
}
