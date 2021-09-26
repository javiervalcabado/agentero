
package policy_data

import (
	"fmt"	
)

type PolicyHolder struct {	// users
	ID			 string		`json:"ID"`
	Name 		 string		`json:"name"`
	MobileNumber string 	`json:"mobile_number"`	
}

func (policyHolder *PolicyHolder) PrintInfo () {
	fmt.Println("Policy Holder " + policyHolder.Name + " with ID " + policyHolder.ID + " (mobile phone: " + policyHolder.MobileNumber + ")")
}