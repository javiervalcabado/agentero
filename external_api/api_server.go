
package external_api

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
	"errors"

	"agentero/policy_data" 			// Data structures regarding policies
)

var (
	filePath = "external_api/data.json"

	allData = AllData{}
)

type AllData struct {
	AllUsers	[]policy_data.PolicyHolder		`json:"users"`
	AllPolicies	[]policy_data.InsurancePolicy	`json:"policies"`
	AllAgents	[]policy_data.InsuranceAgent 	`json:"agents"`
}

// Struct used to manage the gRPC methods called by client
type ContactAndPolicies struct {
	User 		policy_data.PolicyHolder 		`json: "user"`
	Policies 	[]policy_data.InsurancePolicy	`json: "policies"`
}


func PrepareAPI (url string, port string) (result []byte) {
	fmt.Println("Accesing external API...")

	jsonFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
    	log.Fatal("Error opening API data archive: " + err. Error())
	}
	fmt.Println("Accesing data archive in external API .json")

	// Saving our opened jsonFile as a byte array.
	jsonAsBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error reading API data archive: " + err.Error())
	}

	// Closing file once we read it and saved the data
	defer jsonFile.Close()

	// We need to check if all mobile fones have unneeded characters, and are only numbers
	// For that, from []bytes we convert to struct AllData, access the values and work with them
	json.Unmarshal(jsonAsBytes, &allData)	
	fixMobileNumbers()

	fmt.Println("API data: updated ")

	return result
}

func fixMobileNumbers() (){

	var policies = []policy_data.InsurancePolicy{} 

	for i:=0; i<len(allData.AllPolicies); i++ {
		policy := allData.AllPolicies[i]

		// Fix mobile Number: if not a valid number, discard policy
		mobileNumber, err := policy_data.CheckMobileNumber(policy.MobileNumber)
    	if err != nil {
    		fmt.Println("Ignoring policy entry: " + err.Error())
    	} else {
    		policy.MobileNumber = mobileNumber
    		policies = append(policies, policy)
    	}
	}
	allData.AllPolicies = policies
}

// GET/contact/policies/:userID in a real HTTP call
func ContactAndPoliciesById (ID string) (contactAndPolicies ContactAndPolicies, err error) {

	// First we get the User by its ID
	for i:=0; i<len(allData.AllUsers); i++ {
		user := allData.AllUsers[i]
		if user.ID == ID {
			contactAndPolicies.User = user
			// When we get the correct user, we get the associated policies
			contactAndPolicies.Policies = UserPolicies(user)
			return contactAndPolicies, nil

		}
	}
	// If the user hasn't been found, return error
	return ContactAndPolicies{}, errors.New("User with ID " + ID + " not found")	
}

// GET/contact/policies/mp/:userMobileNumber in a real HTTP call
func ContactAndPoliciesByMobileNumber (mobileNumber string) (contactAndPolicies ContactAndPolicies, err error) {
	// First we get the User by its mobile number
	for i:=0; i<len(allData.AllUsers); i++ {
		user := allData.AllUsers[i]
		if user.MobileNumber == mobileNumber {
			contactAndPolicies.User = user
			// When we get the correct user, we get the associated policies
			contactAndPolicies.Policies = UserPolicies(user)
			return contactAndPolicies, nil

		}
	}
	// If the user hasn't been found, return error
	return ContactAndPolicies{}, errors.New("User with mobile number " + mobileNumber + " not found")
}


// Returns policies asociated with an user. The relation is the mobileNumber
func UserPolicies (user policy_data.PolicyHolder) (result []policy_data.InsurancePolicy) {
	
	for i:=0; i<len(allData.AllPolicies); i++ {
		policy := allData.AllPolicies[i]

		if (user.MobileNumber == policy.MobileNumber){
			result = append(result, policy)
		}
	}
	return result
}