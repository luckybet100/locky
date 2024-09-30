package main

import (
	pb "github.com/Locky-Inc/locky/internal/generated/protobuf/terminal"
	"github.com/Locky-Inc/locky/protobuf/terminal"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":4343")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGatewayServer(s, &terminal.Handler{})
	log.Println("listening...")
	log.Fatal(s.Serve(lis))
}