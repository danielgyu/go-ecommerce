package order

import (
	"context"
	"flag"
	"log"
	"net/http"
	"strings"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"
	tl "github.com/danielgyu/go-ecommerce/pkg/tools"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

func allHandlers(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func RunOrderServer() {
	p := flag.String("port", ":8002", "binded port number")

	grpcOrder := grpc.NewServer()
	db := NewMysqlClient()
	pb.RegisterOrderServcieServer(grpcOrder, NewOrderHandler(db))

	httpMux := runtime.NewServeMux()
	err := pb.RegisterOrderServcieHandlerFromEndpoint(context.Background(), httpMux, *p, []grpc.DialOption{grpc.WithInsecure()})
	tl.Check(err)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("healthy")) })

	log.Printf("Running grpc server at port %s\n", *p)
	err = http.ListenAndServe(*p, allHandlers(grpcOrder, mux))
	tl.Check(err)
}
