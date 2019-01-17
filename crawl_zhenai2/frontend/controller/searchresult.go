package controller

import (
	"github.com/olivere/elastic"
	"net/http"
	"strings"
	"strconv"
	"crawl_zhenai2/frontend/view"
	"crawl_zhenai2/frontend/model"
	"context"
	"fmt"
	"reflect"
	"crawl_zhenai2/engine"
	"regexp"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	//client.Index().
	//	Index("datint_profile").
	//	Do(context.Background())
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

//localhost:9527/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	//fmt.Fprintf(w, "q=%s, from=%d", q, from)

	fmt.Printf("q:%s, form:%d\n", q, from)
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.view.Render(w, page)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult

	result.Query = q

	q = rewriteQueryString(q)
	fmt.Println(q)
	resp, err := h.client.Search("datint_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from

	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	//for _, v := range result.Items {
	//	fmt.Printf("%+v\n", v)
	//}

	/*
	if resp.Hits.TotalHits >0{
		for _,hit:=range resp.Hits.Hits{

			var item model2.Profile
			err :=json.Unmarshal(*hit.Source,&item)
			if err != nil{
				panic(err)
			}
			fmt.Printf("%s\n",*hit.Source)
			fmt.Printf("%+v\n",item)
			result.Items = append(result.Items,item)
		}
	}
	*/

	/*
	itemRaw :=resp.Each(reflect.TypeOf(engine.Item{}))
	fmt.Println("len-->",len(itemRaw))
	for _, v := range itemRaw {
		item := v.(engine.Item)
		//fmt.Printf("%+v\n", v)
		result.Items = append(result.Items, item)
	}

*/

	return result, nil

}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
