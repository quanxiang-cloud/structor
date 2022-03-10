package gateway

import (
	"net/http"
	"strings"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

type GatewayServer struct {
	mux  *http.ServeMux
	grpc *grpc.Server
}

func New() *GatewayServer {
	return &GatewayServer{
		mux: http.NewServeMux(),
	}
}

func (g *GatewayServer) Run(port string) error {
	return http.ListenAndServe(port, handler(g.grpc, g.mux))
}

func (g *GatewayServer) RegisterHTTPHandler(path string, f func(w http.ResponseWriter, r *http.Request)) {
	g.mux.HandleFunc(path, f)
}

func (g *GatewayServer) RegisterRpcServer(s *grpc.Server) {
	g.grpc = s
}

func handler(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
