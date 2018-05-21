package wechat

import "net/http"

var client = &http.Client{}

//type AppResponse struct {
//
//}

func get(url string){
	resp, err := client.Get(url)

}