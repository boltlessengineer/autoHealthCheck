package profile

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	. "github.com/gangjun06/auto-selfcheck"
)

type Stdnt struct {
	Name string
	Birth string
	Area Area
	Level Level
	SchNm string
}

func ReadCsv(filePath string) []Stdnt {
	file, _ := os.Open(filePath)

	rdr := csv.NewReader(bufio.NewReader(file))

	rows, _ := rdr.ReadAll()

	stdnts := make([]Stdnt, 0)

	for _, row := range rows[1:] {
		name := row[5]
		birth := row[6]
		area := findAreaNum(row[7])
		level := findLevelNum(row[8])
		schNm := row[9]

		fmt.Printf("%s %s %d %d %s \n", name, birth, area, level, schNm)

		stdnt := Stdnt{name, birth, area, level, schNm}
		stdnts = append(stdnts, stdnt)
		
		for j := 5; j < 10; j++ {
			fmt.Printf("%s ", row[j])
		}

		fmt.Println()
	}

	return stdnts
}

func findAreaNum(areaName string) Area {
	switch areaName {
	case "서울특별시" :
		return 1
	case "부산광역시" :
		return 2
	case "대구광역시" :
		return 3
	case "인천광역시" :
		return 4
	case "광주광역시" :
		return 5
	case "대전광역시" :
		return 6
	case "울산광역시" :
		return 7
	case "세종특별자치시" :
		return 8
	case "경기도" :
		return 9
	case "강원도" :
		return 10
	case "충청북도" :
		return 11
	case "충청남도" :
		return 12
	case "전라북도" :
		return 13
	case "전라남도" :
		return 14
	case "경상북도" :
		return 15
	case "경상남도" :
		return 16
	case "제주특별자치도" :
		return 17
	default :
		return 9
	}
}

func findLevelNum(level string) Level {
	switch level {
	case "초등학교":
		return 2
	case "중학교":
		return 3
	case "고등학교":
		return 4
	default :
		return 4
	}
}
