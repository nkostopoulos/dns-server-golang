package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main(){
	fmt.Println("We are going to parse ifconfig output")

	// Execute "ifconfig" shell command and get the output
	out, err := exec.Command("ifconfig").Output()

	if err != nil {
		log.Fatal(err)
	}

	// Convert output to string
	out2 := string(out)

	// With the following, we can print special characters of a string
	//fmt.Printf("%+q\n", out2)

	// Split the output based on \n\n to get each interface separately
	splitOutSlice := strings.Split(out2, "\n\n")
	splitOutSlice = splitOutSlice[:len(splitOutSlice) - 1]

	var ifaceSlice []map[string]string

	for _, splitOut := range(splitOutSlice) {
		trimmedSplitOut := trimAllSpace(splitOut)
		parts := strings.Fields(trimmedSplitOut)
		var ifaceMap = make(map[string]string)
		for index, word := range(parts) {
			if index == 0 {
				ifaceMap["iface"] = word[:len(word) - 1]
			} 
			if word == "inet" {
				ifaceMap["inet"] = parts[index + 1]
			} else if word == "netmask" {
				ifaceMap["netmask"] = parts[index + 1]
			} else if word == "broadcast" {
				ifaceMap["broadcast"] = parts[index + 1]
			}
		}
		fmt.Println(ifaceMap)
		ifaceSlice = append(ifaceSlice, ifaceMap)
	}
	fmt.Println("-----------------")
	fmt.Println(ifaceSlice)
}

func trimAllSpace(s string) string {
    return strings.Join(strings.Fields(s), " ")
}