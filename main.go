package main

import (
	"bufio"
	"fmt"
	"os"

	checker "github.com/seongmin8452/2020/autoHealthCheck/autoChecker"
	profile "github.com/seongmin8452/2020/autoHealthCheck/readProfile"
)

func main() {
	usr := profile.ReadProfile("./profile/profile.json")
	fmt.Println("user name   :", usr.Name)
	fmt.Println("user birth  :", usr.Birth)
	fmt.Println("user shcool :", usr.School)
	fmt.Println("\n--Auto Student Health Check Start!--")
	rtnMsg := checker.Autocheck(usr)
	fmt.Println("------------------------------------")
	fmt.Println(rtnMsg)
	fmt.Println("------------------------------------")
	fmt.Println("done.")
	fmt.Println("\n[ Press 'Enter' to exit ]")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
