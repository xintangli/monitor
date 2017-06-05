package main

import (
	"fmt"
	"strconv"
	"encoding/json"
	"time"
)

type TestStruct struct {
	SEcho int `json:"sEcho"`
	ITotalRecords int `json:"iTotalRecords"`
	ITotalDisplayRecords int `json:"iTotalDisplayRecords"`
	AaData []string `json:"aaData"`
	AaDataValue []int `json:"aaDataValue"`
}

func main() {
	testInterval()
}

func testInterval()  {
	start, _ := time.Parse("2006-01-02 03:04","2017-05-20 10:00")
	end, _ := time.Parse("2006-01-02 03:04","2017-05-20 10:10")

	interval := end.Unix() - start.Unix()

	if interval > 10 * 60 {
		fmt.Println(interval)
	}
}


func testList()  {
	str := "{\"sEcho\": 3,\"iTotalRecords\": 57,\"iTotalDisplayRecords\": 57,\"aaData\": [\"a\",\"b\",\"c\"],\"aaDataValue\" : [1,2,4]}"
	testStruc := new(TestStruct)
	json.Unmarshal([]byte(str), &testStruc)
	fmt.Println(testStruc)
}

func testFormatFloat()  {
	f := 1.20
	fmt.Println(f)
	var s string
	s = strconv.FormatFloat(f, 'f', 3, 64)
	fmt.Println(s)
}
