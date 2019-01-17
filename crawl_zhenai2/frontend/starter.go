package main

import (
	"net/http"
	"crawl_zhenai2/frontend/controller"
)

func main() {
	http.Handle("/",http.FileServer(http.Dir("frontend/view")))


	http.Handle("/search",controller.CreateSearchResultHandler("frontend/view/template.html"))
	err := http.ListenAndServe(":9527",nil)
	if err != nil{
		panic(err)
	}
}
