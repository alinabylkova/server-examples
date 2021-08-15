package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// grpcShopService := service.New(db)

	grpcServer := grpc.NewServer()

	// gen.RegisterShopServiceServer(grpcServer, &grpcShopService)

	// start the server
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	stopped := make(chan struct{})
	t := time.NewTimer(time.Second * 5)
	go func() {
		grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-t.C:

		log.Print("graceful grpc server shutdown timed out, forcing shutdown")
		grpcServer.Stop()
	case <-stopped:
		t.Stop()
	}
}
