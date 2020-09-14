package checker

import (
	"fmt"
	"log"
	"net/http"

	selfcheck "github.com/gangjun06/auto-selfcheck"
)

func Autocheck() string {
	orgCode, err := selfcheck.FindSchool("보평고등학교", 9, 4)
	checkErr(err)
	fmt.Println(orgCode)

	info, err := selfcheck.GetStudnetInfo(9, orgCode, "이성민", "030801")
	checkErr(err)

	err = info.AllHealthy()
	checkErr(err)
	return "asdgho"
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
