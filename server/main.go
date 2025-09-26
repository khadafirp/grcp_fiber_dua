package server

import (
	"context"
	"log"
	"net"

	"github.com/khadafirp/grpc_fiber_dua/grpc_fiber_dua/proto/greeter"

	"google.golang.org/grpc"
)

type barangServer struct {
	greeter.UnimplementedBarangServiceServer
}

func (s *barangServer) AllBarang(ctx context.Context, req *greeter.BarangRequest) (*greeter.BarangReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &greeter.BarangReply{Message: "Hello, " + req.GetName()}, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	greeter.RegisterBarangServiceServer(grpcServer, &barangServer{})

	log.Println("gRPC server listening on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
