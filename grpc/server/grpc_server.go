package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	pb "github.com/sanghvisagar/go-server/cmd/grpc/proto"
	"google.golang.org/grpc"
)

type XyzServer struct {
	pb.UnimplementedMyServiceServer
}

func main() {

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
	listener, _ := net.Listen("tcp", ":1923")

	grpcServer := grpc.NewServer()

	pb.RegisterMyServiceServer(grpcServer, &XyzServer{})

	grpcServer.Serve(listener)

}

func (UserMgmtServer *XyzServer) MyMethod(ctx context.Context, myRequest *pb.MyRequest) (*pb.MyResponse, error) {
	fmt.Println("Hello", myRequest.Name)
	return &pb.MyResponse{Message: "sagar"}, nil
}
