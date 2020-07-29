package main

import (
	"fmt"

	checker "github.com/seongmin8452/2020/autoHealthCheck/autoChecker"
)

func main() {
	fmt.Println("Auto Student Health Check Start!")
	checker.Autocheck()
	fmt.Println("done.")
}
