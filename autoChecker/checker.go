package checker

import (
	"fmt"

	. "github.com/gangjun06/auto-selfcheck"
	. "github.com/boltlessengineer/autoHealthCheck/readProfile"
)

func Autocheck(s Stdnt) string {
	orgCode, err := FindSchool(s.SchNm, s.Area, s.Level)
	checkErr(err)

	info, err := GetStudnetInfo(9, orgCode, s.Name, s.Birth)
	checkErr(err)

	err = info.AllHealthy()
	checkErr(err)
	return "자가진단이 모두 정상으로 완료되었습니다."
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
