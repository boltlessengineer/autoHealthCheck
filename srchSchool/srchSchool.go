package srchSchool

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Json response struct
type ResultSVO struct {
	RtnRsltCode string
	SchulCode   string
	SchulNm     string
}

type ResponseTemplate struct {
	ResultSVO ResultSVO
}

func SrchSchool(srchNm string) (string, string, string) {
	return getSchoolCode(srchNm)
}

func getSchoolCode(srchNm string) (string, string, string) {
	var baseURL string = "https://eduro.goe.go.kr/stv_cvd_co00_004.do"

	res, err := http.PostForm(baseURL, url.Values{
		"schulNm": {srchNm},
	})
	checkErr(err)
	checkCode(res)

	body, err := ioutil.ReadAll(res.Body)
	checkErr(err)
	// fmt.Println(string(body))

	defer res.Body.Close()

	data := new(ResponseTemplate)
	json.Unmarshal(body, &data)

	RtnRsltCode := data.ResultSVO.RtnRsltCode
	SchulCode := data.ResultSVO.SchulCode
	SchulNm := data.ResultSVO.SchulNm

	return RtnRsltCode, SchulCode, SchulNm
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
