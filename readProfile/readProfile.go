package profile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/seongmin8452/2020/autoHealthCheck/srchSchool"
)

type Profile struct {
	Student Student
}

type Student struct {
	Name       string
	Birth      string
	School     string
	SchoolCode string
}

func ReadProfile(filePath string) *Student {
	if !fileExists(filePath) {
		fmt.Println("[!] Profile.json file doesn't exists")
		makeUsrProfile(filePath)
	}
	file, _ := ioutil.ReadFile(filePath)
	fmt.Println(string(file))
	data := new(Profile)
	err := json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return &data.Student
}

func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func makeUsrProfile(filePath string) {
	var name, birth, prevSchoolNm, schoolNm, schoolCode string
	fmt.Println("------------------------------------------------")
	fmt.Printf("학생 이름 입력                     : ")
	fmt.Scan(&name)
	fmt.Printf("생년월일(주민번호 첫 6자리) 입력   : ")
	fmt.Scan(&birth)

	// /*
	for {
		var usrSelect string = "y"
		fmt.Printf("학교 이름 입력                     : ")
		fmt.Scan(&prevSchoolNm)
		_, schoolCode, schoolNm = srchSchool.SrchSchool(prevSchoolNm)
		if prevSchoolNm != schoolNm {
			fmt.Println("학교명이 입력한 것과 다릅니다.")
			fmt.Printf(schoolNm)
			fmt.Printf("(이)가 맞습니까? (Y/n) : ")
			fmt.Scan(&usrSelect)
			if strings.ToLower(usrSelect) == "y" {
				break
			}
		} else {
			break
		}
	}
	// */

	defaultProfile := []byte("{\n    \"student\": {\n        \"name\": \"" + name + "\",\n        \"birth\": \"" + birth + "\",\n        \"school\": \"" + schoolNm + "\",\n        \"schoolcode\": \"" + schoolCode + "\"\n    }\n}\n")

	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	_, err2 := f.Write(defaultProfile)
	if err2 != nil {
		log.Fatalln(err2)
	}
	defer f.Close()
}
