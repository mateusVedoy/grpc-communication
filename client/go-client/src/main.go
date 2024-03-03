package main

import "github.com/golesson/go-grpc-client/src/api"

func main() {
	route := api.NewRouter()
	route.Start()
}
