package main

import (
	"fmt"
	"log"
	"net"

	pb "./agentero"
)

const (
	address     = "localhost:50051"
	defaultName = "Valkaaaaa"
)

type server struct {
	// Method created in agenter_grpc.pb.go
	pb.UnimplementedCommsServer
}

func main() {
    fmt.Println("Client online")

    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	greeterClient := pb.NewCommsClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := greeterClient.CredentialSystem(ctx, &pb.LogRequest{
		Name: name,
		Pass: "huehuehue"
		})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	if r.GetSuccess() {
		log.Printf("(whipers) I'm in")	
	}

	








    //fmt.Scanln() // wait for Enter Key
}

