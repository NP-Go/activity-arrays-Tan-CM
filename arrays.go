package main

import "fmt"

func main() {
	//Insert Code here
	var a [10]string

	a[0] = "Operating System List"
	a[1] = "Windows XP"
	a[2] = "Linux 16.04"
	a[3] = "Raspbian Teensy"
	a[4] = "MacOS"
	a[5] = "IOS"
	a[6] = "Google Android"

	for i := 0; i < len(a); i++ {
		fmt.Printf(a[i]+" = %d\n", len(a[i]))
	}
}
