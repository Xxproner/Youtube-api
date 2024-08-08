package main

import "github.com/subchen/go-curl"

func main() {
	var Rq = curl.NewRequest()

	Rq.AddCookie(cookie)
}
