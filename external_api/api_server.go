
package external_api

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
)

var (
	filePath = "external_api/data.json"
)

func ReadData () (result []byte) {
	fmt.Println("Accesing external API...")

	jsonFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
    	log.Fatal("Error opening data archive in 'external' API: " + err. Error())
	}
	fmt.Println("Accesing data archive in external API .json")

	// Saving our opened jsonFile as a byte array.
	jsonAsBytes, _ := ioutil.ReadAll(jsonFile)

	// Closing file once we read it and saved the data
	jsonFile.Close()

	return jsonAsBytes
}