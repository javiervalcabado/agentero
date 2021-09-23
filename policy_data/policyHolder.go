
package policy_data

import (
	"fmt"	
)

type PolicyHolder struct {	// users
	Name 		 string		`json:"name"`
	MobileNumber string 	`json:"mobile_number"`	
}

func (policyHolder *PolicyHolder) PrintPolicyHolder () {
	fmt.Println("Policy Holder " + policyHolder.Name + " (mobile phone: " + policyHolder.MobileNumber + ")")
}