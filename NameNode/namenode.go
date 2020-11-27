package main

import (
	"fmt"
	"log"
	"net"

	"./ClienteName"
	"google.golang.org/grpc"
)

func main() {

	/*_, err := os.Create("./LOG.txt")
	if err != nil {
		log.Fatal(err)
	}*/

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := ClienteName.Server{}

	grpcServer := grpc.NewServer()

	ClienteName.RegisterNameServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
