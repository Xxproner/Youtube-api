package main

import (
	"fmt"

	"github.com/subchen/go-curl"
)

func main() {
	var Rq = curl.NewRequest()
	Rq.URL = "https://www.youtube.com/youtubei/v1/backstage/create_post?prettyPrint=false"
	fmt.Println(Rq)
	//Rq.AddCookie(cookie)
}
