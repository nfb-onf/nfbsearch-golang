package controllers

import "fmt"
import "github.com/robfig/revel"
import "github.com/mattbaird/elastigo/api"


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

	var uriVal string
	uriVal = fmt.Sprintf("/%s/%s/_search", "nfb_films", "films")
	body, _ := api.DoCommand("POST", uriVal, nil, query)

	return c.RenderText(string(body[:]))
}
