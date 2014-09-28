package service

import "github.com/gohttp/app"
import "github.com/gohttp/response"
import "github.com/gohttp/logger"
import elastigo "github.com/mattbaird/elastigo/lib"
import "log"
import "net/http"

// import "encoding/json"
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
	// response.JSON(res, pkgs)
}

func (s *Service) SearchHandler(res http.ResponseWriter, req *http.Request) {
	p := params(res, req)

	//TODO: clean this up with query dsl
	search, err := s.Db.Search("versions", "packages", nil, map[string]interface{}{
		"query": map[string]interface{}{
			"filtered": map[string]interface{}{
				"query": map[string]interface{}{
					"match_all": map[string]string{},
				},
				"filter": map[string]interface{}{
					"and": []map[string]interface{}{
						{
							"term": map[string]string{
								"name": p["name"],
							},
						},
						{
							"term": map[string]string{
								"type": p["type"],
							},
						},
					},
				},
			},
		},
	})

	if err != nil {
		u.Error(res, err)
		return
	}

	j := []interface{}{}

	for _, h := range search.Hits.Hits {
		j = append(j, h.Source)
	}

	response.JSON(res, j)
}

func (s *Service) CreateHandler(res http.ResponseWriter, req *http.Request) {
	p := params(res, req)

	s.Db.Index("versions", "packages", "", nil, Package{
		Name: p["name"],
		Url:  p["url"],
		Type: p["type"],
	})

	res.WriteHeader(http.StatusCreated)
}

func (s *Service) ShowHandler(res http.ResponseWriter, req *http.Request) {
}

func params(res http.ResponseWriter, req *http.Request) map[string]string {
	return map[string]string{
		"type": req.URL.Query().Get(":type"),
		"name": req.URL.Query().Get("name"),
	}
}

func (s *Service) Init() {
	s.Use(logger.New())
	s.Get("/", http.HandlerFunc(s.HomeHandler))
	s.Get("/types/:type/packages", http.HandlerFunc(s.IndexHandler))
	s.Get("/types/:type/packages/search", http.HandlerFunc(s.SearchHandler))
	s.Get("/types/:type/packages/:name", http.HandlerFunc(s.ShowHandler))
	s.Post("/types/:type/packages", http.HandlerFunc(s.CreateHandler))
}
