package main

import (
	"context"
	"flag"

	"github.com/quanxiang-cloud/structor/api/rpc"
)

func main() {
	var port string
	flag.StringVar(&port, "port", ":80", "grpc port")

	flag.Parse()

	ctx := context.Background()
	server, err := rpc.New(ctx)
	if err != nil {
		panic(err)
	}

	server.Run(port)
}
