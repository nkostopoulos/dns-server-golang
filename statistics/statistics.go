package statistics

import (
	"fmt"
)

type Stats struct {
	name string
	frequency uint
}

var statsTableSlice = []Stats{}

func CountResolutions(name string) {
	var found bool = false
	for index, statsTable := range(statsTableSlice) {
		if statsTable.name == name {
			statsTable.frequency = statsTable.frequency + 1
			statsTableSlice[index] = statsTable
			found = true

			break
		}
	}

	if !found {
		statsTable := Stats{} 
		statsTable.name = name
		statsTable.frequency = 1
		statsTableSlice = append(statsTableSlice, statsTable)
	}
	fmt.Println(statsTableSlice)
}