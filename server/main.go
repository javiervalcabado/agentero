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
	port = ":8080"
)

var (
	// Flag for runtime (go run <this file's route> -schedule_period=<minutes>)
	// -schedule_period = --schedule_period
	schedule_period = flag.Int("schedule_period", 0, "Minutes between each server call afer initial call")
	apiData 		= APIData{}
	users 			= []policy_data.PolicyHolder{}
	policies 		= []policy_data.InsurancePolicy{}
)

type server struct {
	// Method created in agenter_grpc.pb.go
	pb.UnimplementedCommsServer
}


// The data we are receiving from the 'external API'
type APIData struct {
	Users 		[]policy_data.PolicyHolder		`json: "users"`
	Policies 	[]policy_data.InsurancePolicy	`json: "policies"`
}




func main() {
    fmt.Println("Starting server...")
  

  	importData()

    //startServer()




    fmt.Scanln() // wait for Enter Key
}


func startServer () {
    fmt.Println("Starting server...")

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

	    fmt.Println("Continuing with  server...")

}

// The server receives the request from the client and simulates a credential check
func (server *server) CredentialSystem (ctx context.Context, logRequest *pb.LogRequest) (*pb.LogReply, error) {
	log.Printf("Received log request from %v (password=%v)", logRequest.GetName(), logRequest.GetPass())
	return &pb.LogReply{
		Success: true,
	}, nil
}


// Receives all API data and populates the local DDBB with that information
// Upgrade: checks if the data is already in memory so there are no repetitions
func importData () {
	

	flag.Parse()
	fmt.Print("Flag 'schedule_period' = " + strconv.Itoa(*schedule_period) + ", so we call the API ")

  	if (*schedule_period == 0) {
  		fmt.Println("just once")
  		callAPI()
  	} else {
  		fmt.Println("every " +  strconv.Itoa(*schedule_period) + " minutes")
	  	scheduler := gocron.NewScheduler(time.UTC)
		scheduler.Every(*schedule_period).Second().Do( 
			func (){
				callAPI()
			})
		scheduler.StartAsync()
  	}



    // Mobile numbers test
    
  	if (1==0) {
  		    mobileNumber, err := policy_data.CheckMobileNumber("1111111111")
    if err != nil {
    	fmt.Println("Ignoring policy entry: " + err.Error())
    }

    mobileNumber, err = policy_data.CheckMobileNumber("2(222)222222")
    if err != nil {
    	fmt.Println("Ignoring policy entry: " + err.Error())
    }

    mobileNumber, err = policy_data.CheckMobileNumber("3-3(3)33-33333")
    if err != nil {
    	fmt.Println("Ignoring policy entry: " + err.Error())
    }

    mobileNumber, err = policy_data.CheckMobileNumber("444444444")
    if err != nil {
    	fmt.Println("Ignoring policy entry: " + err.Error())
    }    

    fmt.Println(mobileNumber)
  	}



}

// Calls external API and stores the contents of the database
func callAPI () {
	fmt.Println("Calling external API...")
	

	if (1 == 0) {
		jsonContents := api.ReadData()

	apiData := APIData{}
	json.Unmarshal(jsonContents, &apiData)
	fmt.Println(apiData.Users[1])

	fmt.Println("********")

	fmt.Println(apiData.Policies[1])		
	}
	

	//for i := 0; i < len(users.Users); i++ {
	//}
}



// Returns a policy holder and its policies by userID
func GetContactAndPoliciesById (userID int) {

}

// Returns a single policyholder and its policies by MobileNumber
func GetContactsAndPoliciesByMobileNumber (mobileNumber string) {

}
