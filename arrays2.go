package main

import "fmt"

func main() {
	//Insert Code here
	//var a [10]string

	// Using Slice
	//slice := make ([] type, <length>, ,<capacity>)
	slic := []string{
		"Operating System List",
		"Windows XP",
		"Linux 16.04",
		"Raspbian Teensy",
		"MacOS",
		"IOS",
		"Google Android",
	}

	for i := 0; i < len(slic); i++ {
		fmt.Printf(slic[i]+" = %d\n", len(slic[i]))
	}
}
