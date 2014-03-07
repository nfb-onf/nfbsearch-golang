package controllers

import "github.com/robfig/revel"
import "github.com/mattbaird/elastigo/api"
import "github.com/mattbaird/elastigo/core"


type Payload struct {
	Key string
	Value string
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Search(q string) revel.Result {
	api.Domain = "vmdev.nfb.ca"
	api.Port = "9200"

	var query = map[string]interface {}{
		"query": map[string]interface {}{
			"bool": map[string]interface {}{
				"should": map[string]interface {}{
					"query_string": map[string]interface {} {
						"default_field": "_all",
						"query": q,
					},
				},
			},
		},
	}

	results, error := core.SearchRequest("nfb_films", "films", nil, query)
	return c.RenderJson(map[string]interface {}{
		"results": results,
		"error": error,
	})
}
