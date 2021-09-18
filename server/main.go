package main

import (
	"fmt"
	"flag"  	// Needed for flags on startup
	"net"		// Needed for connection client-server
	"context"
	pb 		"../agentero" //"Protocol Buffer"
)

var (
	// Flag for runtime (go run <this file's route> -schedule_period=<minutes>)
	// -schedule_period = --schedule_period (one or two '-' mean the same)
	schedule_period = flag.Int("schedule_period", false, "Minutes between each server call afer initial call")
	port      		= flag.Int("port", 8080, "Server port")
)

type server struct {
	// Method created in agenter_grpc.pb.go
	pb.UnimplementedCommsServer
}

func main() {
    fmt.Println("Server online")
    startServer()
    // Initial call
/*
    if (schedule_period = 0) {
    }
    */
    //fmt.Scanln() // wait for Enter Key
}


func startServer () {
	flag.Parse()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		  log.Fatalf("Error listening port %d: %v", port, err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

// The server receives the request from the client and simulates a credential check
func (server *server) LogRequest (ctx context.Context, helloRequest *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received log request from %v (password=%v)", helloRequest.GetName(), helloRequest.GetPass())
	return &pb.LogReply{
		Name: 	 helloRequest.GetName(),
		Success: true 
	}

}

func (server *server) LogReply