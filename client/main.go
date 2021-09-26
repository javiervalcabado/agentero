package main

import (
	"fmt"
	"log"
	//"os"
	"context"
	"time"
	"encoding/json"

	"google.golang.org/grpc"

	pb "agentero/agentero"	// Protocol Buffer

	"agentero/policy_data" // Data structures regarding policies
)


const (
	address     = "localhost:8080"
	defaultName = "Guy Incognito"
)

// Struct that stores the result of gRPC methods GetContactAndPoliciesById or GetContactsAndPoliciesByMobileNumber 
type ContactAndPolicies struct {
	User		policy_data.PolicyHolder		`json:"user"`
	Policies 	[]policy_data.InsurancePolicy	`json:"policies"`
} 


func main() {
    fmt.Println("Starting client...")
   
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	commsClient := pb.NewCommsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	// Simple gRPC method to check connectivity
	fmt.Println("Sending credentials to log system...")
	gRPC1, err := commsClient.CredentialSystem(ctx, &pb.LogRequest{
		Name: defaultName,
		Pass: "blablabla",
	})
	if err != nil {
		log.Fatalf("gRPC1 failed: %v", err)
	}
	if gRPC1.GetSuccess() {
		fmt.Println("---- Succesful login ----")	
	} else {
		log.Fatalf("---- Login didn't work ----")
	}

	fmt.Println("---------------------------------------------------------------------")

	// Local structs whre we store the information we ask the server
	contactAndPolicies1 := ContactAndPolicies{}
	contactAndPolicies2 := ContactAndPolicies{}

	// gRPC method that ask for a single policy holder and its policies by userID.
	fmt.Println("Asking server for policy holder and its policies by userID...")
	gRPC2, err := commsClient.GetContactAndPoliciesById(ctx, &pb.UserID {
		ID: "1",
	})
	if err != nil {
		log.Fatalf("gRPC2 failed: %v", err)
	}
	if gRPC2.GetSuccess() {
		fmt.Println("-- User found --")	

		json.Unmarshal(gRPC2.GetContent(), &contactAndPolicies1)	

		contactAndPolicies1.User.PrintInfo()

		for i := 0; i < len(contactAndPolicies1.Policies); i++ {
			contactAndPolicies1.Policies[i].PrintInfo()
		}
	} else {
		fmt.Println("-- User not found --")
	}

	fmt.Println("---------------------------------------------------------------------")

	// gRPC method that ask for a single policy holder and its policies by mobileNumber
	fmt.Println("Asking server for policy holder and its policies by user mobile phone...")
	// TODO: quitar este bucle para poder acceder al código
	gRPC3, err := commsClient.GetContactsAndPoliciesByMobileNumber(ctx, &pb.MobileNumber {
		Number: "1234567892",
	})
	if err != nil {
		log.Fatalf("gRPC3 failed: %v", err)
	}
	if gRPC3.GetSuccess() {
		fmt.Println("-- User found --")	

		json.Unmarshal(gRPC3.GetContent(), &contactAndPolicies2)	

		contactAndPolicies2.User.PrintInfo()

		for i := 0; i < len(contactAndPolicies2.Policies); i++ {
			contactAndPolicies2.Policies[i].PrintInfo()
		}
	}
	
	fmt.Println("End of client code")
}

