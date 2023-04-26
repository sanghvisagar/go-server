package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/sanghvisagar/go-server/cmd/grpc/proto"
	"google.golang.org/grpc"
)

type XyzServer struct {
	pb.UnimplementedMyServiceServer
}

func main() {
	listener, _ := net.Listen("tcp", ":1923")

	grpcServer := grpc.NewServer()

	pb.RegisterMyServiceServer(grpcServer, &XyzServer{})

	grpcServer.Serve(listener)

}

func (UserMgmtServer *XyzServer) MyMethod(ctx context.Context, myRequest *pb.MyRequest) (*pb.MyResponse, error) {
	fmt.Println("Hello", myRequest.Name)
	return &pb.MyResponse{Message: "sagar"}, nil
}
