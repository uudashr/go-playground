package main

import "net/http"

func main() {
	svr := &http.Server{}
	if err := svr.ListenAndServe(); err != nil {
		panic(err)
	}
}
