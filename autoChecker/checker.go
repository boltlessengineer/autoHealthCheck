package checker

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	profile "github.com/seongmin8452/2020/autoHealthCheck/readProfile"
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
func Autocheck(user *profile.Student) string {
	var baseURL string = "https://eduro.goe.go.kr/stv_cvd_co02_000.do"
	qstnCrtfcNoEncpt, rtnRsltCode := getSVO(user)

	res, err := http.PostForm(baseURL, url.Values{
		"qstnCrtfcNoEncpt": {qstnCrtfcNoEncpt},
		"rtnRsltCode":      {rtnRsltCode},
		"schulNm":          {user.School},
		"pName":            {user.Name},
		"frnoRidno":        {user.Birth},
		"schulCode":        {user.SchoolCode},
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
	//fmt.Println(string(body))

	defer res.Body.Close()

	s1 := strings.Split(string(body), "<P style=\"margin:5rem auto; line-height:3rem; text-align:center;\">")[1]
	content_text := strings.Trim(strings.Split(s1, "</p>")[0], " ")
	textLines := strings.Split(content_text, "<br>")
	textLines[0] = strings.TrimSpace(textLines[0])[6:]
	textLines[1] = strings.Replace(strings.TrimSpace(textLines[1]), "<br/>", "\n", 1)

	rtnMsg := strings.Join(textLines[:], "\n")

	return rtnMsg
}

func getSVO(user *profile.Student) (string, string) {
	var baseURL string = "https://eduro.goe.go.kr/stv_cvd_co00_012.do"

	res, err := http.PostForm(baseURL, url.Values{
		"schulNm":   {user.School},
		"pName":     {user.Name},
		"frnoRidno": {user.Birth},
		"schulCode": {user.SchoolCode},
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

func getRtnMsg(resBody io.Reader) string {
	doc, err := goquery.NewDocumentFromReader(resBody)
	checkErr(err)
	_, b := doc.Attr(".content_box")
	if !b {
		fmt.Println(doc.Text())
	}

	content_box := doc.Find(".content_box").Text()

	return content_box
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
