package service

import "github.com/gohttp/app"
import "github.com/gohttp/response"
import "github.com/gohttp/logger"
import elastigo "github.com/mattbaird/elastigo/lib"
import "net/http"
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

func (s *Service) IndexHandler(res http.ResponseWriter, req *http.Request) {
	p := params(res, req)

	a := []map[string]interface{}{}

	a = append(a, map[string]interface{}{
		"match": map[string]string{
			"name": p["type"],
		},
	})

	search, err := s.Db.Search("registry", "packages", nil, map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": a,
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

func (s *Service) SearchHandler(res http.ResponseWriter, req *http.Request) {
	p := params(res, req)

	if p["name"] == "" {
		s.IndexHandler(res, req)
		return
	}

	a := []map[string]interface{}{}

	a = append(a, map[string]interface{}{
		"match": map[string]string{
			"name": p["name"],
		},
	})

	a = append(a, map[string]interface{}{
		"match": map[string]string{
			"type": p["type"],
		},
	})

	search, err := s.Db.Search("registry", "packages", nil, map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": a,
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

	s.Db.Index("registry", "packages", "", nil, Package{
		Name: p["name"],
		Url:  p["url"],
		Type: p["type"],
	})

	res.WriteHeader(http.StatusCreated)
}

func (s *Service) ShowHandler(res http.ResponseWriter, req *http.Request) {
	p := params(res, req)

	result, err := s.search(p["name"], p["type"])

	if err != nil {
		u.Error(res, err)
		return
	}

	if len(result.Hits.Hits) > 0 {
		response.JSON(res, result.Hits.Hits[0].Source)
		return
	} else {
		res.WriteHeader(http.StatusNotFound)
	}
}

func (s *Service) search(name string, t string) (elastigo.SearchResult, error) {
	return s.Db.Search("registry", "packages", nil, map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					map[string]interface{}{
						"match": map[string]string{
							"name.untouched": name,
						},
					},
					map[string]interface{}{
						"match": map[string]string{
							"name": t,
						},
					},
				},
			},
		},
	})
}

func (s *Service) DeleteHandler(res http.ResponseWriter, req *http.Request) {
	p := params(res, req)

	result, err := s.search(p["name"], p["type"])

	if err != nil {
		u.Error(res, err)
		return
	}

	for _, h := range result.Hits.Hits {
		_, err := s.Db.Delete("registry", "packages", h.Id, nil)

		if err != nil {
			u.Error(res, err)
			return
		}
	}

	res.WriteHeader(http.StatusNoContent)
}

func params(res http.ResponseWriter, req *http.Request) map[string]string {
	name := req.URL.Query().Get("name")
	if name == "" {
		name = req.URL.Query().Get(":name")
	}
	return map[string]string{
		"type": req.URL.Query().Get(":type"),
		"name": name,
		"url":  req.URL.Query().Get("url"),
	}
}

func (s *Service) Init() {
	s.Use(logger.New())
	s.Get("/types/:type/packages", http.HandlerFunc(s.IndexHandler))
	s.Get("/types/:type/packages/search", http.HandlerFunc(s.SearchHandler))
	s.Get("/types/:type/packages/:name", http.HandlerFunc(s.ShowHandler))
	s.Del("/types/:type/packages/:name", http.HandlerFunc(s.DeleteHandler))
	s.Post("/types/:type/packages", http.HandlerFunc(s.CreateHandler))
}
