package zonefile

import (
	"fmt"
	"os"
	"log"
	"encoding/csv"
)

func LoadNamesFromFile(filename string) []map[string]string {
	// Load Resource Records from the CSV file
	data := openFile(filename)
	resourceRecords := make([]map[string]string, 0)
	for _, line := range(data){
		var name string = line[0]
		var ip string = line[1]
		var tempResourceRecord = make(map[string]string)
		tempResourceRecord["name"] = name
		tempResourceRecord["ip"] = ip
		resourceRecords = append(resourceRecords, tempResourceRecord)
	}

	return resourceRecords
}

func openFile(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func PrintResourceRecords(resourceRecords []map[string]string){
	// Print all the loaded Resource Records
	for _, mapItem := range(resourceRecords) {
		fmt.Printf("The name is %v and the IP is %v\n", mapItem["name"], mapItem["ip"])
	}
	fmt.Println()
}

func ResolveName(resourceRecords []map[string]string) map[string]string {
	var name string
	fmt.Println("Provide a DNS that you want to resolve")
	fmt.Scan(&name)

	var ipAddress string = "None"
	for _, mapItem := range(resourceRecords) {
		var mapName string = mapItem["name"]
		if mapName == name {
			ipAddress = mapItem["ip"]
		}
	}

	if ipAddress == "None" {
		fmt.Printf("No known IP address for name %v\n", name)
	} else {
		fmt.Printf("The IP address corresponding to name %v is %v\n", name, ipAddress)
	}

	var resolutionData = make(map[string]string)
	resolutionData["name"] = name
	resolutionData["ipAddress"] = ipAddress
	return resolutionData
}