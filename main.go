package main

import (
	accountRpc "github.com/zbwang163/ad_content_overpass"
	"github.com/zbwang163/ad_content_server/common/client"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	client.Init()
	lis, err := net.Listen("tcp", ":50002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	accountRpc.RegisterContentServiceServer(s, NewServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
