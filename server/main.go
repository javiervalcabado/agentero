package main

import (
	"fmt"
	"flag"  	// Needed for flags on startup
	"net"		// Needed for connection client-server
	"log"
	"context"
	"google.golang.org/grpc"

	pb 	"agentero/agentero" //"Protocol Buffer"
)

const (
	port = ":8080"
)

var (
	// Flag for runtime (go run <this file's route> -schedule_period=<minutes>)
	// -schedule_period = --schedule_period (one or two '-' mean the same)
	schedule_period = flag.Int("schedule_period", 0, "Minutes between each server call afer initial call")
)

type server struct {
	// Method created in agenter_grpc.pb.go
	pb.UnimplementedCommsServer
}

func main() {
    fmt.Println("Starting server...")
    startServer()
    // Initial call
/*
    if (schedule_period = 0) {
    }
    */
    //fmt.Scanln() // wait for Enter Key
}


func startServer () {
	//flag.Parse()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		  log.Fatalf("Error listening port %d: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCommsServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error serving to serve: %v", err)
	}


}

// The server receives the request from the client and simulates a credential check
func (server *server) CredentialSystem (ctx context.Context, logRequest *pb.LogRequest) (*pb.LogReply, error) {
	log.Printf("Received log request from %v (password=%v)", logRequest.GetName(), logRequest.GetPass())
	return &pb.LogReply{
		Success: true,
	}, nil
}

