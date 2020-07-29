package main

import (
	"fmt"

	checker "github.com/seongmin8452/2020/autoHealthCheck/autoChecker"
	profile "github.com/seongmin8452/2020/autoHealthCheck/readProfile"
)

func main() {
	usr := profile.ReadProfile("profile.json")
	fmt.Println("user name   :", usr.Name)
	fmt.Println("user birth  :", usr.Birth)
	fmt.Println("user shcool :", usr.School)
	fmt.Println("\n--Auto Student Health Check Start!--")
	rtnMsg := checker.Autocheck(usr)
	fmt.Println("------------------------------------")
	fmt.Println(rtnMsg)
	fmt.Println("------------------------------------")
	fmt.Println("done.")
	fmt.Println("\n[ Press any key to exit ]")
	var input string
	fmt.Scan(&input)
}
