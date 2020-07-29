package checker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Json response Struct from 012.do
type ResultSVO struct {
	RtnRsltCode      string
	QstnCrtfcNoEncpt string
}

type ResponseFrom12 struct {
	ResultSVO ResultSVO
}

// AutoCheck automatically checks student-health-check
func Autocheck() {
	var baseURL string = "https://eduro.goe.go.kr/stv_cvd_co02_000.do"
	qstnCrtfcNoEncpt, rtnRsltCode := getSVO()

	res, err := http.PostForm(baseURL, url.Values{
		"qstnCrtfcNoEncpt": {qstnCrtfcNoEncpt},
		"rtnRsltCode":      {rtnRsltCode},
		"schulNm":          {"보평고등학교"},
		"pName":            {"이성민"},
		"frnoRidno":        {"030801"},
		"schulCode":        {"J100005836"},
		"rspns01":          {"1"},
		"rspns02":          {"1"},
		"rspns07":          {"0"},
		"rspns08":          {"0"},
		"rspns09":          {"0"},
	})
	checkErr(err)
	checkCode(res)

	body, err := ioutil.ReadAll(res.Body)
	checkErr(err)
	fmt.Println(string(body))

	defer res.Body.Close()
}

func getSVO() (string, string) {
	var baseURL string = "https://eduro.goe.go.kr/stv_cvd_co00_012.do"

	res, err := http.PostForm(baseURL, url.Values{
		"schulNm":   {"보평고등학교"},
		"pName":     {"이성민"},
		"frnoRidno": {"030801"},
		"schulCode": {"J100005836"},
	})
	checkErr(err)
	checkCode(res)

	body, err := ioutil.ReadAll(res.Body)
	checkErr(err)

	defer res.Body.Close()

	data := new(ResponseFrom12)
	json.Unmarshal(body, &data)
	RtnRsltCode := data.ResultSVO.RtnRsltCode
	QstnCrtfcNoEncpt := data.ResultSVO.QstnCrtfcNoEncpt
	fmt.Println(QstnCrtfcNoEncpt)

	return QstnCrtfcNoEncpt, RtnRsltCode
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
