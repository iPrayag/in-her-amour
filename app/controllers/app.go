package controllers

import "github.com/revel/revel"
import "fmt"
import "log"
//import "encoding/json"
import "github.com/mattbaird/elastigo/api"
import "github.com/mattbaird/elastigo/core"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Feeling(feel string) revel.Result {
  GetFeels()
	return c.Render(feel)
}

//@see : http://www.arkxu.com/post/58405378919/play-elasticsearch-with-go
//@see : https://github.com/mattbaird/elastigo/blob/master/client.go

func GetFeels(){

	api.Domain = "localhost"
 
	// search users
	var searchQuery = `{
		"query": {
			"term": {"firstName":"J"}
		}
	}`
 
	response, err := core.SearchRequest("gccount", "Customer", nil, searchQuery)
	if err != nil {
		log.Fatalf("The search of users has failed:", err)
	}
//	var values []interface{}

	for _, hit := range response.Hits.Hits {
/*
		var value map[string]interface{}
		err := json.Unmarshal(hit.Source, &value)
		if err != nil {
			log.Fatalf("Failed to unmarshal, line 43", err)
		}
		values = append(values, value)
*/
    fmt.Println(hit)
	}
	//fmt.Println(values)

}
