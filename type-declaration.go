package main

import "fmt"

func main(){
	type noKTP string
	type Married bool

	var noKtpSofi noKTP = "3512156804000001"
	var marriedStatus Married = true
	fmt.Println(noKtpSofi)
	fmt.Println(marriedStatus)
}