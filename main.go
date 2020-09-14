package main

import (
	"fmt"

	checker "github.com/seongmin8452/2020/autoHealthCheck/autoChecker"
)

func main() {
	fmt.Println("-------Auto Student Health Checker v0.9.2-------")
	//usr := profile.ReadProfile("profile.json")
	fmt.Println("------------------------------------------------")
	//fmt.Println("user name   :", usr.Name)
	//fmt.Println("user birth  :", usr.Birth)
	//fmt.Println("user school :", usr.School)
	fmt.Println("\n--------Auto Student Health Check Start!--------")
	rtnMsg := checker.Autocheck()
	fmt.Println("------------------------------------------------")
	fmt.Println(rtnMsg)
	fmt.Println("------------------------------------------------")
	fmt.Println("done.")
	fmt.Println("\n[ Press any key to exit ]")
	var input string
	fmt.Scan(&input)
}
