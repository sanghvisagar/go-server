package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/sanghvisagar/go-server/cmd/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:1923", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, _ := c.MyMethod(ctx, &pb.MyRequest{Name: "sagar"})

	fmt.Println("Hello ", resp.GetMessage())

}
