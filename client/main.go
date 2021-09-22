package main

import (
	"fmt"
	"log"
	//"os"
	"context"
	"time"

	"google.golang.org/grpc"

	pb "agentero/agentero"	// Protocol Buffer

	//"agentero/policy_data" // Data structures regarding policies
)


const (
	address     = "localhost:8080"
	defaultName = "Guy Incognito"
)



func main() {
    fmt.Println("Starting client...")
   
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	commsClient := pb.NewCommsClient(conn)

/*
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
*/

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	r, err := commsClient.CredentialSystem(ctx, &pb.LogRequest{
		Name: defaultName,
		Pass: "blablabla",
		})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.GetSuccess() {
		log.Printf("-- Succesful login --")	
	} else {
		log.Printf("-- Login didn't work --")
	}

    fmt.Scanln() // wait for Enter Key
}

