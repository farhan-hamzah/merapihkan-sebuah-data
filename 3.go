package main

import (
	"fmt"
)
func inpData(nama *string, status *string, angkatan *int) {
	fmt.Scan(nama, status, angkatan)
}
func outputData(nama string, status string, angkatan int) {
	fmt.Printf("| %-11s | %-11s | %-8d |\n", nama, status, angkatan)
}

func main() {
	var n, i, angkatan int
	var nama, status string
	i = 0
	fmt.Scan(&n)
	for i< n {
		inpData(&nama, &status, &angkatan)
		if i == 0 {
			fmt.Println("----------------------------------------")
			fmt.Println("|                DATA                  |")
			fmt.Println("----------------------------------------")
		}
		outputData(nama, status, angkatan)
		i++
	}
	fmt.Println("----------------------------------------")
}