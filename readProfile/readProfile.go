package profile

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	file, _ := ioutil.ReadFile(filePath)
	data := new(Profile)
	err := json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return &data.Student
}
