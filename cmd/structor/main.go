package main

import (
	"context"
	"flag"

	"github.com/quanxiang-cloud/structor/api/rpc"
)

func main() {
	var port string
	var suffix string
	flag.StringVar(&port, "port", ":80", "grpc port")
	flag.StringVar(&suffix, "suffix", "clause", "suffix")

	flag.Parse()

	ctx := context.Background()
	server, err := rpc.New(ctx, rpc.WithSuffix(suffix))
	if err != nil {
		panic(err)
	}

	server.Run(port)
}
