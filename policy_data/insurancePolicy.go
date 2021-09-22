
package policy_data

import (
	"fmt"
	"strings"
	"errors"
	"strconv"
)

type InsurancePolicy struct {
	MobileNumber 	string		`json: "mobile_number"`
	Premium 		string		`json: "premium"`
	Type 			string 		`json: "type"`
}

func (insurancePolicy *InsurancePolicy) PrintInsurancePolicy () {
	fmt.Println("Policy mobile number: " + insurancePolicy.MobileNumber + ", Premium: " + insurancePolicy.Premium + ", Type: " + insurancePolicy.Type)
}

// Deletes '(', ')', '-', and extra espaces 
func CheckMobileNumber (mobileNumberString string) (string, error) {
	mobileNumber  := strings.Replace(mobileNumberString, "(", "", -1)
	//fmt.Println("1: " + mobileNumber)
	mobileNumber  = strings.Replace(mobileNumber, ")", "", -1)
	//fmt.Println("2: " + mobileNumber)
	mobileNumber  = strings.Replace(mobileNumber, "-", "", -1)
	//fmt.Println("3: " + mobileNumber)
	mobileNumber  = strings.Replace(mobileNumber, " ", "", -1)
	//fmt.Println("3: " + mobileNumber)

	_, err := strconv.Atoi(mobileNumber)
	if err != nil {
		fmt.Println("Error checking telephone number " + mobileNumber + " (NaN)")
		return mobileNumber, err
	}

	if len(mobileNumber) != 10 {
		fmt.Println("Error checking telephone number " + mobileNumber + " (length != 10)")	
		return mobileNumber, errors.New("Length of mobile number not equals to 10")
	}

	return mobileNumber, nil
}

