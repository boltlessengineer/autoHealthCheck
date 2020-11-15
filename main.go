package main

import (
	"fmt"
	. "github.com/boltlessengineer/autoHealthCheck/readProfile"
	checker "github.com/boltlessengineer/autoHealthCheck/autoChecker"
)

func main() {
	data := ReadCsv("./data.csv")
	fmt.Println("-------Auto Student Health Checker v0.9.2-------")
	fmt.Println("------------------------------------------------")
	fmt.Println("\n--------Auto Student Health Check Start!--------")
	for _, stdnt := range data {
		rtnMsg := checker.Autocheck(stdnt)
		fmt.Println(rtnMsg)
		fmt.Println("------------------------------------------------")
	}
	fmt.Println("done.")
	fmt.Println("\n[ Press any key to exit ]")
	var input string
	fmt.Scan(&input)
}
