package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/quanxiang-cloud/cabin/logger"
	"github.com/quanxiang-cloud/structor/api/gateway"
	"github.com/quanxiang-cloud/structor/api/rpc"
	"github.com/quanxiang-cloud/structor/pkg/probe"
)

var (
	port string

	level           int
	development     bool
	initial         int
	thereafter      int
	outputPath      VarStringArray
	errorOutputPath VarStringArray
)

func main() {
	catchVar()

	log := logger.New(&logger.Config{
		Level:       level,
		Development: development,
		Sampling: logger.Sampling{
			Initial:    initial,
			Thereafter: thereafter,
		},
		OutputPath:      outputPath,
		ErrorOutputPath: errorOutputPath,
	})

	server := gateway.New()

	ctx := context.Background()
	grpcServer, err := rpc.New(ctx)
	if err != nil {
		panic(err)
	}

	probe := probe.New(log)

	server.RegisterRpcServer(grpcServer)
	server.RegisterHTTPHandler("/liveness", probe.LivenessProbe)
	server.RegisterHTTPHandler("/readiness", probe.ReadinessProbe)

	server.Run(port)
}

type VarStringArray []string

func (v *VarStringArray) String() string {
	return fmt.Sprint(*v)
}

func (v *VarStringArray) Set(s string) error {
	*v = append(*v, s)
	return nil
}

func catchVar() {
	flag.StringVar(&port, "port", ":80", "grpc port")
	flag.IntVar(&level, "level", -1, "log level")
	flag.BoolVar(&development, "development", false, "log development")
	flag.IntVar(&initial, "initial", 100, "log initial")
	flag.IntVar(&thereafter, "thereafter", 100, "log thereafter")
	flag.Var(&outputPath, "outputPath", "log outputPath")
	flag.Var(&errorOutputPath, "errorOutputPath", "log errorOutputPath")

	flag.Parse()

	if len(outputPath) == 0 {
		outputPath.Set("stderr")
	}

	if len(errorOutputPath) == 0 {
		errorOutputPath.Set("stderr")
	}
}
