package main

import (
	"fmt"
	"dns-server/zonefile"
	"dns-server/statistics"
)

func main(){
	var filename string = "./zonefile.csv"
	resourceRecords := zonefile.LoadNamesFromFile(filename)

	fmt.Println("The resource records loaded from the zone file are the following:")
	zonefile.PrintResourceRecords(resourceRecords)

	var resolutionData = make(map[string]string)
	for {
		resolutionData = zonefile.ResolveName(resourceRecords)
		statistics.CountResolutions(resolutionData["name"])
	}
}