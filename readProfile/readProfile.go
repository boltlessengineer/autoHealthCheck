package profile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
		fmt.Println("Profile.json file doesn't exists")
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
	var name, birth string
	fmt.Printf("학생 이름 입력                   : ")
	fmt.Scan(&name)
	fmt.Printf("생년월일(주민번호 첫 6자리) 입력 : ")
	fmt.Scan(&birth)
	defaultProfile := []byte("{\n    \"student\": {\n        \"name\": \"" + name + "\",\n        \"birth\": \"" + birth + "\",\n        \"school\": \"보평고등학교\",\n        \"schoolcode\": \"J100005836\"\n    }\n}\n")

	//err := ioutil.WriteFile(filePath, defaultProfile, 0644)
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
