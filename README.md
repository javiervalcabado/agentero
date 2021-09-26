# agentero


---------------------------------------
-									  -	 
- Agentero Golang Backend assignment  - 
-									  -
---------------------------------------


1. Start server/main.go 
	Starts a gRPC server to listen for client/main.go in localhost:8080
	Also prepares the data in external_api/api_server
	Optional flags:
		a) schedule_period: set the timer in minutes for a refresh of the API data. Default is 0, which only creates that data once
		b) ams-api-url: could be used to connect to a real API, doesn't do anything now

3. Start client/main.go
	It connects to the gRPC server set by server/main.go
	It sends a mock login for testing purposes
	It asks for gRPC methods GetContactAndPoliciesById("1") and GetContactsAndPoliciesByMobileNumber("1234567892") for this example

-----------------------------------------------------------------------------------------------------------------------------------------------------
