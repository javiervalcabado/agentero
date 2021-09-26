package main

import (
	"fmt"
	"flag" 
	"net"
	"log"
	"context"
	"time"
	"encoding/json"
	"strconv"
	
	"google.golang.org/grpc"
	"github.com/go-co-op/gocron"

	"agentero/policy_data" 			// Data structures regarding policies
	pb 	"agentero/agentero" 		//"Protocol Buffer"
	api "agentero/external_api"		// Calling the 'external' API and managing its data
)

const (
	gRPCPort = ":8080"
	APIPort	 = ":8081"
)

var (
	// Flag for runtime (go run <this file's route> -schedule_period=<minutes>)
	// -schedule_period = --schedule_period
	schedule_period = flag.Int("schedule_period", 0, "Minutes between each server call afer initial call")

	// Flag for AMS API calls. Right now it's useless
	ams_api_url = flag.String("ams-api-url", "", "URL for connection to AMS API")
	
	// Struct to unmarshal API data
	apiData 		= APIData{}
	jsonContents	= []byte{}
	
	// Structures that store locally the data from API calls
	users 			= []policy_data.PolicyHolder{}
	policies 		= []policy_data.InsurancePolicy{}
)

type server struct {
	// Method created in agenter_grpc.pb.go
	pb.UnimplementedCommsServer
}


// Not used right now, could store all different information from the API or a subset
type APIData struct {
	Users 		[]policy_data.PolicyHolder		`json: "users"`
	Policies 	[]policy_data.InsurancePolicy	`json: "policies"`
}

// Struct used to manage the gRPC methods called by client
type ContactAndPolicies struct {
	User 		policy_data.PolicyHolder 		`json: "user"`
	Policies 	[]policy_data.InsurancePolicy	`json: "policies"`
}

func main() {
    fmt.Println("Starting...")
  
  	importData()

    startServer()

}

// We launch the gRPC server to listen for the client
func startServer () {
    fmt.Println("Starting server...")

	lis, err := net.Listen("tcp", gRPCPort)
	if err != nil {
		  log.Fatalf("Error listening port %d: %v", gRPCPort, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCommsServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error serving to serve: %v", err)
	}

	fmt.Println("Continuing with server...")
}

// The server receives the request from the client and simulates a credential check
func (server *server) CredentialSystem (ctx context.Context, logRequest *pb.LogRequest) (*pb.LogReply, error) {
	log.Println("Received log request from " + logRequest.GetName() + " (password=" + logRequest.GetPass() + ")")
	return &pb.LogReply{
		Success: true,
	}, nil
}

// The server receives the request for the user and its policies by its ID
func (server *server) GetContactAndPoliciesById (ctx context.Context, userID *pb.UserID) (*pb.ContactAndPolicies, error) {
	fmt.Println("Received contact & policies request (userID = " + userID.GetID() + ")")

	// We ask the API for the data and then we prepare it to send: if anything fails, it returns error
	apiResult, err := api.ContactAndPoliciesById(userID.GetID())
	result, err := json.Marshal(apiResult)

	if err != nil {
		return &pb.ContactAndPolicies {
		Success: false,
		Content: nil,
		}, err
	} 

	// If everything went right, we send the resut from the API to the client
	return &pb.ContactAndPolicies {
		Success: true,
		Content: result,
	}, nil
}



//  The server receives the request for the user and its policies by its mobile number
func (server *server) GetContactsAndPoliciesByMobileNumber (ctx context.Context, userMobileNumber *pb.MobileNumber) (*pb.ContactAndPolicies, error) {
	fmt.Println("Received contact & policies request (mobile number = " + userMobileNumber.GetNumber() + ")")

	// We ask the API for the data and then we prepare it to send: if anything fails, it returns error
	apiResult, err := api.ContactAndPoliciesByMobileNumber(userMobileNumber.GetNumber())
	result, err := json.Marshal(apiResult)

	if err != nil {
		return &pb.ContactAndPolicies {
		Success: false,
		Content: nil,
		}, err
	} 

	// If everything went right, we send the resut from the API to the client
	return &pb.ContactAndPolicies {
		Success: true,
		Content: result,
	}, nil
}


// Calls the external API based on the schedule_period flag
func importData () {
	
	flag.Parse()
	fmt.Print("Flag 'schedule_period' = " + strconv.Itoa(*schedule_period) + ", so we call the API ")

  	if (*schedule_period == 0) {
  		fmt.Println("just once")
  		api.PrepareAPI(*ams_api_url, APIPort)
  	} else {
  		fmt.Println("every " +  strconv.Itoa(*schedule_period) + " minutes")
	  	scheduler := gocron.NewScheduler(time.UTC)
		scheduler.Every(*schedule_period).Minute().Do( 
			func (){
				api.PrepareAPI(*ams_api_url, APIPort)
			})
		scheduler.StartAsync()
  	}
}